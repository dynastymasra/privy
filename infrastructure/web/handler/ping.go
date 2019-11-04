package handler

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/privy/delivery/http/formatter"

	"github.com/dynastymasra/privy/infrastructure/database/postgres"
	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/privy/config"
	"github.com/sirupsen/logrus"
)

func Ping(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithField(config.RequestID, r.Context().Value(config.HeaderRequestID))

		if err := postgres.Ping(db); err != nil {
			log.WithError(err).Errorln("Failed ping postgres")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.SuccessResponse().Stringify())
	}
}
