package server

import (
	"net/http"
	"reviews/errdefs"
	"reviews/http/httputils"
	"reviews/http/middlewares"
	"reviews/http/router"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Logging bool
}

type Server struct {
	cfg         *ServerConfig
	httpServer  *http.Server
	routers     []router.Router
	middlewares []middlewares.MiddleWare
}

func NewServer(cfg *ServerConfig) *Server {
	return &Server{cfg: cfg}
}

func (srv *Server) UseMiddleware(m middlewares.MiddleWare) {
	srv.middlewares = append(srv.middlewares, m)
}

func (srv *Server) UseRouter(routers ...router.Router) {
	srv.routers = append(srv.routers, routers...)
}

func (srv *Server) Accept(addr string) error {
	srv.httpServer = &http.Server{Addr: addr, Handler: srv.createMux()}
	return srv.httpServer.ListenAndServe()
}

func (srv *Server) makeHandler(handlerFunc httputils.ApiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r, mux.Vars(r)); err != nil {
			if errdefs.IsHttpError(err) {
				status := errdefs.GetCode(err)
				http.Error(w, err.Error(), status)
			} else {
				http.Error(w, "system error", 500)
			}
		}
	}
}

func (srv *Server) handleWithMiddlewares(handler httputils.ApiHandlerFunc) httputils.ApiHandlerFunc {
	next := handler

	for _, m := range srv.middlewares {
		next = m.WrapHandler(next)
	}

	if srv.cfg.Logging {
		loggerMiddleware := middlewares.NewLoggerMiddleware([]string{})

		next = loggerMiddleware.WrapHandler(next)
	}

	return next
}

func (srv *Server) createMux() *mux.Router {
	ro := mux.NewRouter()

	for _, router := range srv.routers {
		for _, route := range router.Routes() {
			f := srv.makeHandler(route.Handler())

			ro.Path(route.Path()).Methods(route.Method()).Handler(f)
		}
	}

	return ro
}
