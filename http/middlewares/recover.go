package middlewares

import (
	"net/http"
	"reviews/http/httputils"

	"github.com/sirupsen/logrus"
)

type RecoverMiddleware struct{}

func (r RecoverMiddleware) WrapHandler(handler httputils.ApiHandlerFunc) httputils.ApiHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("Recovered from panic: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()

		return handler(w, r, vars)
	}
}
