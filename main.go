package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dynastymasra/privy/console"
	"github.com/dynastymasra/privy/infrastructure/database/postgres"

	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/privy/config"
	"github.com/urfave/cli"
	"gopkg.in/tylerb/graceful.v1"
)

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Database initialization
	postgresDB, err := postgres.Connect(config.Postgres())
	if err != nil {
		logrus.WithError(err).Fatalln("Unable to open connection to postgres")
	}

	clientApp := cli.NewApp()
	clientApp.Name = config.ServiceName
	clientApp.Version = config.Version

	clientApp.Action = func(c *cli.Context) error {
		server := &graceful.Server{
			Timeout: 0,
		}

		select {
		case sig := <-stop:
			<-server.StopChan()
			logrus.Warningln(fmt.Sprintf("Service shutdown because %+v", sig))

			if err := postgresDB.Close(); err != nil {
				logrus.WithError(err).Errorln("Unable to turn off Postgres connections")
			}

			logrus.Infoln("Postgres Connection closed")
			os.Exit(0)
		}

		return nil
	}

	clientApp.Commands = []cli.Command{
		{
			Name:        "migrate:run",
			Description: "Running Migration",
			Action: func(c *cli.Context) error {
				if err := console.RunDatabaseMigrations(postgresDB.DB()); err != nil {
					os.Exit(1)
				}
				return nil
			},
		},
		{
			Name:        "migrate:rollback",
			Description: "Rollback Migration",
			Action: func(c *cli.Context) error {
				if err := console.RollbackLatestMigration(postgresDB.DB()); err != nil {
					os.Exit(1)
				}
				return nil
			},
		},
		{
			Name:        "migrate:create",
			Description: "Create up and down migration files with timestamp",
			Action: func(c *cli.Context) error {
				return console.CreateMigrationFiles(c.Args().Get(0))
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
