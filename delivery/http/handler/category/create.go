package category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/delivery/http/formatter"
	"github.com/dynastymasra/privy/domain"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type categoryRequest struct {
	Name       string `json:"name" validate:"required,max=255"`
	Enable     bool   `json:"enable" validate:"omitempty"`
	ProductIDs []int  `json:"products,omitempty" validate:"omitempty"`
}

func CreateHandler(service domain.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithFields(logrus.Fields{
			config.RequestID: r.Context().Value(config.HeaderRequestID),
		})

		var reqBody categoryRequest

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorln("Unable to read request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		if err := json.Unmarshal(body, &reqBody); err != nil {
			log.WithError(err).WithField("body", string(body)).Errorln("Unable to parse request body")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		if err := validator.New().Struct(&reqBody); err != nil {
			log.WithError(err).WithField("body", reqBody).Errorln("Failed validate category request")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		cat := domain.Category{
			Name:       reqBody.Name,
			Enable:     reqBody.Enable,
			ProductIDs: reqBody.ProductIDs,
		}

		category, err := service.Create(r.Context(), cat)
		if err != nil {
			log.WithError(err).WithField("body", reqBody).Errorln("Failed create new category")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, formatter.ObjectResponse(category).Stringify())
	}
}
