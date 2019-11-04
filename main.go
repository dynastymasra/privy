package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dynastymasra/privy/image"

	"github.com/dynastymasra/privy/category"

	"github.com/dynastymasra/privy/domain"

	"github.com/dynastymasra/privy/product"

	"github.com/dynastymasra/privy/infrastructure/web"

	"github.com/dynastymasra/privy/console"
	"github.com/dynastymasra/privy/infrastructure/database/postgres"

	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/privy/config"
	"github.com/icrowley/fake"
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

	productRepository := product.NewRepository(postgresDB)
	productService := product.NewService(productRepository)

	categoryRepository := category.NewRepository(postgresDB)
	categoryService := category.NewService(categoryRepository)

	imageRepository := image.NewRepository(postgresDB)
	imageService := image.NewService(imageRepository)

	clientApp := cli.NewApp()
	clientApp.Name = config.ServiceName
	clientApp.Version = config.Version

	clientApp.Action = func(c *cli.Context) error {
		server := &graceful.Server{
			Timeout: 0,
		}
		go web.Run(server, postgresDB, web.ServiceInstance{
			Product:  productService,
			Category: categoryService,
			Image:    imageService,
		})
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
				return console.RunDatabaseMigrations(postgresDB.DB())
			},
		},
		{
			Name:        "migrate:rollback",
			Description: "Rollback Migration",
			Action: func(c *cli.Context) error {
				return console.RollbackLatestMigration(postgresDB.DB())
			},
		},
		{
			Name:        "migrate:create",
			Description: "Create up and down migration files with timestamp",
			Action: func(c *cli.Context) error {
				return console.CreateMigrationFiles(c.Args().Get(0))
			},
		},
		{
			Name:        "migrate:seed",
			Description: "Create up and down migration files with timestamp",
			Action: func(c *cli.Context) error {
				_, err := productRepository.Create(context.Background(), domain.Product{
					Name:        fake.ProductName(),
					Description: fake.Paragraph(),
					Enable:      true,
					Images:      []domain.Image{{ID: 1, Name: fake.ProductName(), File: fake.Paragraphs(), Enable: true}},
					Categories:  []domain.Category{{ID: 1, Name: fake.ProductName(), Enable: true}},
				})

				return err
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
