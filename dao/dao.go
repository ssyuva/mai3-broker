package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mcarloai/mai-v3-broker/common/postgres"
	logger "github.com/sirupsen/logrus"
)

type DAO interface {
	PerpetualDAO
	OrderDAO
	MatchTransactionDAO
	WatcherDAO
	ForUpdate()
	Transaction(body func(DAO) error) error
}

type gormDAO struct {
	db *gorm.DB
	perpetualDAO
	orderDAO
	matchTransactionDAO
	watcherDAO
}

func (g *gormDAO) ForUpdate() {
	g.db = g.db.Set("gorm:query_option", "FOR UPDATE")
}

func (g *gormDAO) Transaction(body func(DAO) error) (err error) {
	tx := g.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Fatal Commit Panic:%s", r)
			logger.Errorf(err.Error())
			tx.Rollback()
		}
	}()

	err = body(NewFromGormDB(tx))

	if err == nil {
		err = tx.Commit().Error
		if err != nil {
			err = fmt.Errorf("Commit:%w", err)
		}
	}

	// Makesure rollback when Block error or Commit error
	if err != nil {
		tx.Rollback()
	}
	return
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
		db:                  db,
		perpetualDAO:        perpetualDAO{db: db},
		orderDAO:            orderDAO{db: db},
		matchTransactionDAO: matchTransactionDAO{db: db},
		watcherDAO:          watcherDAO{db: db},
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
