package reviewsapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reviews/http/server"
	"reviews/internal/app/reviewsapp/movies"
	"reviews/internal/app/reviewsapp/reviews"
	"reviews/internal/app/reviewsapp/routes"
	"reviews/internal/app/reviewsapp/storage/dbstore"
	"reviews/internal/pkg/config"
	"reviews/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type ReviewApp struct {
	Config *config.Config
}

func NewApp() *ReviewApp {
	return &ReviewApp{Config: config.New()}
}

func (ra *ReviewApp) Start() error {
	logger.Init(logger.Config{
		Format: logger.Format(ra.Config.LogFormat),
		Level:  ra.Config.LogLevel,
	})

	srvr := server.NewServer(&server.ServerConfig{Logging: false})
	ctx := context.Background()

	// db, err := reviews.InitDB(reviews.Config{})
	// if err != nil {
	// panic(err)
	// }

	db := &sqlx.DB{}
	movieStore := dbstore.NewMovieStore(ctx, db)

	movieService := movies.NewMovieService(movieStore)
	movierouter := routes.NewMovieRouter(movieService)

	reviewStore := dbstore.NewReviewStore(ctx, db)
	reviewService := reviews.NewReviewService(reviewStore)

	movieRevService := &reviews.MovieReviewService{
		MovieService:  movieService,
		ReviewService: reviewService,
	}

	reviewRouter := routes.NewMovieReviewRouter(movieRevService)

	// srvr.UseMiddleware(middlewares.RecoverMiddleware)
	srvr.UseRouter(routes.NewPingRouter(), movierouter, reviewRouter)

	errCh := make(chan error)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Println("running server at ", ra.Config.Address)
		errCh <- srvr.Accept(ra.Config.Address)
	}()

	select {
	case <-stop:
		return nil
	case errr := <-errCh:
		return errr
	}
}
