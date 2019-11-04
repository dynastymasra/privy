package image

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

func DeleteHandler(service domain.ImageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithFields(logrus.Fields{
			config.HeaderRequestID: r.Context().Value(config.HeaderRequestID),
		})

		v := mux.Vars(r)
		id, err := strconv.Atoi(v["image_id"])
		if err != nil {
			log.WithError(err).Errorln("Failed parse params to int")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		log = log.WithField("image_id", id)

		err = service.Delete(r.Context(), id)
		if err == gorm.ErrRecordNotFound {
			log.WithError(err).Errorln("Failed delete image by id")

			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		if err != nil {
			log.WithError(err).Errorln("Failed delete image by id")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.SuccessResponse().Stringify())
	}
}
