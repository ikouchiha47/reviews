package middlewares

import "reviews/http/httputils"

type MiddleWare interface {
	WrapHandler(httputils.ApiHandlerFunc) httputils.ApiHandlerFunc
}
