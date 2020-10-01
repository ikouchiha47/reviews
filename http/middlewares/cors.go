package middlewares

import (
	"net/http"
	"reviews/http/httputils"
)

type CORSMiddleware struct {
	defaultHeaders string
}

func NewCORSMiddleware(d string) CORSMiddleware {
	return CORSMiddleware{defaultHeaders: d}
}

func (c CORSMiddleware) WrapHandler(handler httputils.ApiHandlerFunc) httputils.ApiHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		corsHeaders := c.defaultHeaders
		if corsHeaders == "" {
			return handler(w, r, vars)
		}

		w.Header().Add("Access-Control-Allow-Origin", corsHeaders)
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth")
		w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, DELETE, PUT, OPTIONS")
		return handler(w, r, vars)
	}
}
