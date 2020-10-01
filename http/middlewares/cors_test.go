package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCorsMiddleware(t *testing.T) {
	corsM := NewCORSMiddleware("*")

	middlewareHandler := func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		originHeader := w.Header().Get("Access-Control-Allow-Origin")

		assert.Equal(t, originHeader, "*")

		return nil
	}

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		corsM.WrapHandler(middlewareHandler)(w, r, map[string]string{})
	}


	handler := http.HandlerFunc(handlerFunc)
	req, err := http.NewRequest(http.MethodGet, "/", nil)

	require.NoError(t, err)
	handler.ServeHTTP(httptest.NewRecorder(), req)
}
