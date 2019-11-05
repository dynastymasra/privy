package product

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

type productRequest struct {
	ID          int    `json:"id" validate:"omitempty"`
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	Enable      bool   `json:"enable" validate:"omitempty"`
	Categories  []int  `json:"categories,omitempty" validate:"omitempty"`
	Images      []int  `json:"images,omitempty" validate:"omitempty"`
}

func CreateHandler(service domain.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		log := logrus.WithFields(logrus.Fields{
			config.RequestID: r.Context().Value(config.HeaderRequestID),
		})

		var reqBody productRequest

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
			log.WithError(err).WithField("body", reqBody).Errorln("Failed validate product request")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		prod := domain.Product{
			ID:          reqBody.ID,
			Name:        reqBody.Name,
			Description: reqBody.Description,
			Enable:      reqBody.Enable,
			CategoryIDs: reqBody.Categories,
			ImageIDs:    reqBody.Images,
		}

		product, err := service.Create(r.Context(), prod)
		if err != nil {
			log.WithError(err).WithField("body", prod).Errorln("Failed create new product")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, formatter.ObjectResponse(product).Stringify())
	}
}
