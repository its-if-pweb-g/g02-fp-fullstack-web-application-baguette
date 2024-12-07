package main

import (
	"api/internal/auth"
	"api/internal/store"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type application struct {
	config        config
	store         store.Storage
	logger        *zap.SugaredLogger
	authenticator auth.Authenticator
}

type config struct {
	addr    string
	db      dbConfig
	apiURL  string
	nextURL string
	auth    authConfig
}

type dbConfig struct {
	addr              string
	maxOpenConnection int
	maxIdleConnection int
	maxIdleTime       string
}

type authConfig struct {
	user   string
	secret string
	exp    time.Duration
	iss    string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware global
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		// Authentication routes
		r.Post("/register", app.registerUserHandler)
		r.Post("/login", app.LoginUserHandler)
		r.Get("/image/{id}", app.ImageHandler)

		// Product routes
		r.Route("/products", func(r chi.Router) {
			r.Get("/", app.ProductsHandler)
			r.Get("/search", app.SearchProductHandler)
			r.Get("/sort", app.SortProductHandler)
		
			r.With(app.AuthTokenMiddleware, app.AdminRoleMiddleware).Post("/", app.CreateProductHandler)
		
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", app.DetailProductHandler)
				r.With(app.AuthTokenMiddleware, app.AdminRoleMiddleware).Put("/", app.UpdateProductHandler)
				r.With(app.AuthTokenMiddleware, app.AdminRoleMiddleware).Delete("/", app.DeleteProductHandler)
			})
		})
				

		// User routes
		r.Route("/user", func(r chi.Router) {
			r.Use(app.AuthTokenMiddleware)

			r.Get("/", app.UserDetailHandler)
			r.Put("/", app.UpdateUserHandler)
			r.Get("/address", app.UserAddressHandler)
			r.Get("/cart", app.UserCartHandler)
			r.Post("/cart/products", app.AddProductToCartHandler) 
			r.Delete("/cart/products/{id}", app.DeleteProductInCartHandler)
			r.Get("/cart/products/inc/{id}", app.IncQuantityHandler)
			r.Get("/cart/products/dec/{id}", app.DecQuantityHandler)
			
			// r.Post("/pay", app.PaymentHandler)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	app.logger.Infow("server has started", "addr", app.config.addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr)

	return nil
}
