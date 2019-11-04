package middleware

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/privy/config"
	"github.com/urfave/negroni"
)

func RequestID() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		requestID := r.Header.Get(config.HeaderRequestID)
		if len(requestID) < 1 {
			requestID = uuid.NewV4().String()

			w.Header().Add(config.HeaderRequestID, requestID)
		}

		ctx := context.WithValue(r.Context(), config.HeaderRequestID, requestID)
		next(w, r.WithContext(ctx))
	}
}
