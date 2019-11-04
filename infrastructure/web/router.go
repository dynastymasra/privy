package web

import (
	"fmt"
	"net/http"

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

	return subRouter
}
