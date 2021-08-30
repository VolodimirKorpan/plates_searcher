package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/VolodimirKorpan/go_kobi/auth"
	authhttp "github.com/VolodimirKorpan/go_kobi/auth/delivery/http"
	authusecase "github.com/VolodimirKorpan/go_kobi/auth/usecase"
	"github.com/VolodimirKorpan/go_kobi/config"
	"github.com/VolodimirKorpan/go_kobi/plates"
	plateshttp "github.com/VolodimirKorpan/go_kobi/plates/delivery/http"
	platesusecase "github.com/VolodimirKorpan/go_kobi/plates/usecase"
	"github.com/VolodimirKorpan/go_kobi/store"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type App struct {
	httpServer *http.Server

	authUC   auth.UseCase
	platesUC plates.UseCase
}

func NewApp() (*App, error) {
	cfg := config.Get()
	ctx := context.Background()

	store, err := store.New(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "store.New failed")
	}

	return &App{
		authUC: authusecase.NewAuthUseCase(
			store.User,
			cfg.HashSalt,
			[]byte(cfg.SigningKey),
			time.Duration(cfg.TokenTTL),
		),
		platesUC: platesusecase.NewPlateUseCase(store.Plate),
	}, nil
}

func (a *App) Run(port string) error {
	//Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//auth endpoints
	authhttp.RegisterHttpEndpoints(router, a.authUC)
	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	plateshttp.RegisterHTTPEndpoints(api, a.platesUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
