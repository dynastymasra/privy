package postgres_test

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/mujib/config"
	"github.com/dynastymasra/mujib/infrastructure/database/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PostgresConnect struct {
	suite.Suite
}

func Test_PostgresConnect(t *testing.T) {
	suite.Run(t, new(PostgresConnect))
}

func (p *PostgresConnect) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (p *PostgresConnect) Test_Connect_Success() {
	db, err := postgres.Connect(config.Postgres())

	assert.NotNil(p.T(), db)
	assert.NoError(p.T(), err)
}

func (p *PostgresConnect) Test_PingDatabase_Success() {
	db, err := postgres.Connect(config.Postgres())

	assert.NoError(p.T(), err)
	assert.NotNil(p.T(), db)

	err = postgres.Ping(db)

	assert.NoError(p.T(), err)

	err = postgres.Close(db)

	assert.NoError(p.T(), err)
}

func (p *PostgresConnect) Test_PingDatabase_Failed() {
	var db *gorm.DB

	err := postgres.Ping(db)

	assert.Error(p.T(), err)
}

func (p *PostgresConnect) Test_CloseDatabase_Failed() {
	var db *gorm.DB

	err := postgres.Close(db)

	assert.Error(p.T(), err)
}

func (p *PostgresConnect) Test_Connect_Failed() {
	viper.Set("DATABASE_HOST", "test")
	config.Load()

	fmt.Println(config.Postgres())

	postgres.Reset()

	db, err := postgres.Connect(config.Postgres())

	assert.Nil(p.T(), db)
	assert.Error(p.T(), err)

	viper.Reset()
	config.Load()
	postgres.Reset()
}
