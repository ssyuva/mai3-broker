package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectPostgres(url string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, fmt.Errorf("models:connect db fail:%w", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(16)

	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}

	return db, nil
}

func IsRecordNotFound(err error) bool {
	var gormError gorm.Errors
	if ok := errors.As(err, &gormError); ok {
		return gorm.IsRecordNotFoundError(gormError)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
