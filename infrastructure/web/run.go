package web

import (
	"net/http"

	"github.com/dynastymasra/privy/image"

	"github.com/dynastymasra/privy/category"

	"github.com/dynastymasra/privy/product"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/privy/config"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
)

type ServiceInstance struct {
	Product  product.Service
	Category category.Service
	Image    image.Service
}

func Run(server *graceful.Server, db *gorm.DB, service ServiceInstance) {
	logrus.Infoln("Start run web application")

	muxRouter := Router(db, service)

	server.Server = &http.Server{
		Addr: config.ServerAddress(),
		Handler: handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true),
			handlers.RecoveryLogger(logrus.StandardLogger()),
		)(muxRouter),
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatalln("Failed to start server")
	}
}
