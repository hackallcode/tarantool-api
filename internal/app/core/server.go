package core

import (
	"net/http"

	"kv-storage/internal/pkg/db"
	"kv-storage/internal/pkg/logger"
	"kv-storage/internal/pkg/middleware"
	"kv-storage/internal/pkg/router"
)

type Params struct {
	Port   string
	Prefix string
}

func StartApp(params Params) error {
	if err := db.Open(); err != nil {
		logger.Error(err.Error())
	}

	r := router.InitRouter(params.Prefix)

	// Middleware
	r.Use(middleware.PanicCatcher)
	r.Use(middleware.Logger)
	r.Use(middleware.ApplyJsonContentType)
	r.Use(middleware.ApplyCors)

	logger.Info("Starting core at " + params.Port)
	return http.ListenAndServe(":"+params.Port, r)
}

func StopApp() {
	logger.Info("Stopping core...")
	if err := db.Close(); err != nil {
		logger.Error(err.Error())
	}
}
