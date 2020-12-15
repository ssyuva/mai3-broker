package dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mcarloai/mai-v3-broker/common/postgres"
	logger "github.com/sirupsen/logrus"
)

type DAO interface {
	PerpetualDAO
	OrderDAO
	MatchTransactionDAO
	LaunchTransactionDAO
	NonceDAO
	KVStoreDAO
	ForUpdate() // mark "ROW SHARE"
	Transaction(ctx context.Context, isReadOnly bool, body func(DAO) error) (err error)
	RepeatableRead(ctx context.Context, isReadOnly bool, body func(DAO) error) (err error)
}

type gormDAO struct {
	db *gorm.DB
	perpetualDAO
	orderDAO
	matchTransactionDAO
	launchTransactionDAO
	nonceDAO
	kvstoreDAO
}

func (g *gormDAO) ForUpdate() {
	newDB := g.db.Set("gorm:query_option", "FOR UPDATE")
	g.db = newDB
	g.perpetualDAO = perpetualDAO{db: newDB}
	g.orderDAO = orderDAO{db: newDB}
	g.matchTransactionDAO = matchTransactionDAO{db: newDB}
	g.launchTransactionDAO = launchTransactionDAO{newDB}
	g.nonceDAO = nonceDAO{newDB}
	g.kvstoreDAO = kvstoreDAO{newDB}
}

func (g *gormDAO) beginTx(ctx context.Context, isolation sql.IsolationLevel, isReadOnly bool, body func(DAO) error) (err error) {
	tx := g.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: isolation,
		ReadOnly:  isReadOnly,
	})
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transaction commit panic:%s", r)
			logger.Errorf(err.Error())
			tx.Rollback()
		}
	}()
	err = body(NewFromGormDB(tx))
	if err == nil {
		err = tx.Commit().Error
		if err != nil {
			err = fmt.Errorf("transaction commit failed:%w", err)
		}
	}

	// Makesure rollback when Block error or Commit error
	if err != nil {
		tx.Rollback()
	}
	return
}

func (g *gormDAO) Transaction(ctx context.Context, isReadOnly bool, body func(DAO) error) (err error) {
	return g.beginTx(context.Background(), sql.LevelReadCommitted, false, body)
}

func (g *gormDAO) RepeatableRead(ctx context.Context, isReadOnly bool, body func(DAO) error) (err error) {
	return g.beginTx(context.Background(), sql.LevelRepeatableRead, isReadOnly, body)
}

var gormDB *gorm.DB

func ConnectPostgres(url string) error {
	db, err := postgres.ConnectPostgres(url)
	if err != nil {
		return err
	}
	gormDB = db
	return nil
}

func IsRecordNotFound(err error) bool {
	return postgres.IsRecordNotFound(err)
}

func New() DAO {
	if gormDB == nil {
		panic("gormDB is uninitialized, connect to database first")
	}
	return NewFromGormDB(gormDB)
}

func NewFromGormDB(db *gorm.DB) DAO {
	return &gormDAO{
		db:                   db,
		perpetualDAO:         perpetualDAO{db: db},
		orderDAO:             orderDAO{db: db},
		matchTransactionDAO:  matchTransactionDAO{db: db},
		launchTransactionDAO: launchTransactionDAO{db},
		nonceDAO:             nonceDAO{db},
		kvstoreDAO:           kvstoreDAO{db},
	}
}

func countWithCap(db *gorm.DB, limit int) (count int64, err error) {
	scope := db.NewScope(db.Value)
	sql := fmt.Sprintf(
		"SELECT COUNT(*) FROM (SELECT * FROM %v %v LIMIT %v) AS count_table",
		scope.QuotedTableName(), scope.CombinedConditionSql(), limit)

	emptyDB := db.New().Model(count)
	err = emptyDB.Raw(sql, scope.SQLVars...).Row().Scan(&count)
	return
}
