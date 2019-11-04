package postgres

import (
	"errors"

	"github.com/dynastymasra/privy/config"
	"github.com/jinzhu/gorm"
	"github.com/matryer/resync"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	err  error
	once resync.Once
)

func Connect(config config.PostgresConfig) (*gorm.DB, error) {
	dbURL := config.ConnectionString()

	once.Do(func() {
		db, err = gorm.Open("postgres", dbURL)
		if err != nil {
			logrus.WithError(err).WithField("db_url", dbURL).Errorln("Cannot connect to DB")
			db = nil
			return
		}

		db.DB().SetMaxIdleConns(config.MaxIdleConn())
		db.DB().SetMaxOpenConns(config.MaxOpenConn())

		if err := db.DB().Ping(); err != nil {
			logrus.WithError(err).Errorln("Cannot ping database")
			db = nil
			return
		}

		db.LogMode(config.LogEnabled())
	})

	return db, err
}

func Ping(db *gorm.DB) error {
	if db == nil {
		return errors.New("does't have database data")
	}
	return db.DB().Ping()
}

func Close(db *gorm.DB) error {
	if db == nil {
		return errors.New("does't have database data")
	}
	return db.DB().Close()
}

func Reset() {
	once.Reset()
}
