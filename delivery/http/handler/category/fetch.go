package category

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/delivery/http/formatter"
	"github.com/dynastymasra/privy/domain"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func FindByIDHandler(service domain.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithFields(logrus.Fields{
			config.HeaderRequestID: r.Context().Value(config.HeaderRequestID),
		})

		v := mux.Vars(r)
		id, err := strconv.Atoi(v["category_id"])
		if err != nil {
			log.WithError(err).Errorln("Failed parse params to int")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		log = log.WithField("category_id", id)

		category, err := service.FindByID(r.Context(), id)
		if err == gorm.ErrRecordNotFound {
			log.WithError(err).Errorln("Failed get category by id")

			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		if err != nil {
			log.WithError(err).Errorln("Failed get category by id")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.ObjectResponse(category).Stringify())
	}
}

func FindAllHandler(service domain.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from := formatter.BuildPaginationFrom(r.FormValue("from"))
		size := formatter.BuildPaginationSize(r.FormValue("size"))

		log := logrus.WithFields(logrus.Fields{
			config.HeaderRequestID: r.Context().Value(config.HeaderRequestID),
			"from":                 from,
			"size":                 size,
		})

		categories, err := service.Fetch(r.Context(), from, size)
		if err != nil {
			log.WithError(err).Errorln("Failed fetch category")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.ObjectResponse(categories).Stringify())
	}
}
