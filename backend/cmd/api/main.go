package main

import (
	"api/internal/auth"
	"api/internal/db"
	"api/internal/env"
	"api/internal/store"
	"context"
	"expvar"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"go.uber.org/zap"
)

func main() {

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Info("No .env file found or error loading it: %v", err)
	}

	cfg := config{
		addr:    env.GetString("GO_API_URL", "localhost:8000"),
		nextURL: env.GetString("NEXT_URL", "http://localhost:3000"),
		db: dbConfig{
			addr:              env.GetString("DATABASE_URL", "mongodb://baguette:bjirlah@localhost:27017"),
			maxOpenConnection: env.GetInt("MONGO_MAX_OPEN_CONNECTION", 30),
			maxIdleConnection: env.GetInt("MONGO_MAX_IDLE_CONNECTION", 30),
			maxIdleTime:       env.GetString("MONGO_MAX_IDLE_TIME", "15m"),
		},
		auth: authConfig{
			user:   env.GetString("AUTH_USER", "admin"),
			secret: env.GetString("AUTH_SECRET", "admin"),
			exp:    time.Hour * time.Duration(env.GetInt("AUTH_EXP", 72)),
			iss:    env.GetString("AUTH_ISS", "baguette"),
		},
		payment: paymentGateway{
			serverKey: env.GetString("SERVER_KEY", ""),
			paymentURL: env.GetString("PAYMENT_URL", "https://app.sandbox.midtrans.com/snap/v1/transactions"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConnection,
		cfg.db.maxIdleConnection,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("database connection pool established")

	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.secret,
		cfg.auth.iss,
		cfg.auth.iss,
	)

	store := store.NewStorage(db)

	addAdmin(store)

	var s snap.Client
	s.New(cfg.payment.serverKey, midtrans.Sandbox)

	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		authenticator: jwtAuthenticator,
	}

	expvar.Publish("database", expvar.Func(func() any {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result := make(map[string]interface{})
		err := db.Database("pweb-g").RunCommand(ctx, map[string]interface{}{"serverStatus": 1}).Decode(&result)
		if err != nil {
			logger.Fatal(err)
			return err
		}
		return result
	}))

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
