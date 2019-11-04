package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/privy/delivery/http/handler/image"

	"github.com/dynastymasra/privy/delivery/http/handler/category"

	"github.com/dynastymasra/privy/delivery/http/handler/product"

	"github.com/dynastymasra/privy/delivery/http/formatter"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/privy/infrastructure/web/middleware"

	"github.com/dynastymasra/privy/infrastructure/web/handler"

	"github.com/dynastymasra/privy/config"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router(db *gorm.DB, service ServiceInstance) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	subRouter := router.PathPrefix("/v1/").Subrouter().UseEncodedPath()
	commonHandlers := negroni.New(
		middleware.RequestID(),
	)

	subRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, formatter.FailResponse(config.ErrDataNotFound.Error()).Stringify())
	})

	subRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, formatter.FailResponse(config.ErrDataNotFound.Error()).Stringify())
	})

	// Probes
	subRouter.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(handler.Ping(db)),
	)).Methods(http.MethodGet, http.MethodHead)

	// product group
	subRouter.Handle("/products", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(product.CreateHandler(service.Product)),
	)).Methods(http.MethodPost)

	subRouter.Handle("/products/{product_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(product.FindByIDHandler(service.Product)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/products", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(product.FindAllHandler(service.Product)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/products/{product_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(product.UpdateHandler(service.Product)),
	)).Methods(http.MethodPut)

	subRouter.Handle("/products/{product_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(product.DeleteHandler(service.Product)),
	)).Methods(http.MethodDelete)

	// category group
	subRouter.Handle("/categories", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(category.CreateHandler(service.Category)),
	)).Methods(http.MethodPost)

	subRouter.Handle("/categories/{category_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(category.FindByIDHandler(service.Category)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/categories", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(category.FindAllHandler(service.Category)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/categories/{category_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(category.UpdateHandler(service.Category)),
	)).Methods(http.MethodPut)

	subRouter.Handle("/categories/{category_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(category.DeleteHandler(service.Category)),
	)).Methods(http.MethodDelete)

	// category group
	subRouter.Handle("/images", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(image.CreateHandler(service.Image)),
	)).Methods(http.MethodPost)

	subRouter.Handle("/images/{image_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(image.FindByIDHandler(service.Image)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/images", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(image.FindAllHandler(service.Image)),
	)).Methods(http.MethodGet)

	subRouter.Handle("/images/{image_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(image.UpdateHandler(service.Image)),
	)).Methods(http.MethodPut)

	subRouter.Handle("/images/{image_id}", commonHandlers.With(
		middleware.HTTPStatLogger(),
		negroni.WrapFunc(image.DeleteHandler(service.Image)),
	)).Methods(http.MethodDelete)

	return subRouter
}
