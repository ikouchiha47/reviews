package middlewares

import (
	"net/http"
	"reviews/http/httputils"
	"reviews/pkg/context"
	"time"

	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type LoggerMiddleware struct {
	ignorePaths []string
}

func NewLoggerMiddleware(ignorePaths []string) LoggerMiddleware {
	return LoggerMiddleware{ignorePaths: ignorePaths}
}

func (l LoggerMiddleware) WrapHandler(handler httputils.ApiHandlerFunc) httputils.ApiHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		reqID := uuid.NewRandom()

		entry := logrus.WithFields(logrus.Fields{
			"reqID": reqID.String(),
		})

		ctx := context.WithReqID(r.Context(), reqID)
		ctx = context.WithLogger(ctx, entry)
		r = r.WithContext(ctx)

		w.Header().Set("X-Request-Id", reqID.String())

		res, ok := w.(negroni.ResponseWriter)
		if !ok {
			res = negroni.NewResponseWriter(w)
		}

		begin := time.Now()

		entry.WithFields(logrus.Fields{
			"StartTime":  begin.Format(time.RFC3339),
			"Status":     res.Status(),
			"DurationMS": time.Since(begin).Milliseconds(),
			"Hostname":   r.Host,
			"Method":     r.Method,
			"Path":       r.URL.Path,
		}).Info("inbound http request")

		return handler(res, r, vars)

	}
}
