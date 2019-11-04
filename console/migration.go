package console

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dynastymasra/mujib/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationSourcePath = "file://migration"
	migrationFilePath   = "./migration"
)

func CreateMigrationFiles(filename string) error {
	if len(filename) == 0 {
		return errors.New("migration filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migrationFilePath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migrationFilePath, timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}
	log.Println("created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	log.Println("created %s\n", downMigrationFilePath)

	return nil
}

func RunDatabaseMigrations(db *sql.DB) error {
	m, err := getMigrate(db)

	if err != nil {
		log.Println(err.Error())
		return joinErrors(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Println("error", err.Error())
		return joinErrors(err)
	}

	log.Println("Migrations successful")
	return nil
}

func RollbackLatestMigration(db *sql.DB) error {
	m, err := getMigrate(db)
	if err != nil {
		return joinErrors(err)
	}
	stepError := m.Steps(-1)
	if stepError != nil {
		log.Println("Rollback error", stepError)
		return joinErrors(stepError)
	}
	log.Println("Migrations Rollback successful")
	return nil
}

func joinErrors(inputErrors error) error {
	var errorMsgs []string
	errorMsgs = append(errorMsgs, inputErrors.Error())
	errMsgJoined := strings.Join(errorMsgs, ",")
	return fmt.Errorf(errMsgJoined)
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func getMigrate(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Println("error", err.Error())
		return nil, err
	}
	return migrate.NewWithDatabaseInstance(
		migrationSourcePath, config.Postgres().Name(), driver)
}
