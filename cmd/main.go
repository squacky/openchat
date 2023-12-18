package main

import (
	"context"
	"log"

	"github.com/squacky/openchat/cmd/httpapi"
	"github.com/squacky/openchat/config"
	"github.com/squacky/openchat/internal/user"
	"github.com/squacky/openchat/pkg/logger"
	"github.com/squacky/openchat/pkg/mongodb"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("starting the message")

	mCfg := mongodb.NewConfig(cfg)
	db, err := mongodb.Connect(context.Background(), mCfg)

	if err != nil {
		logger.Error("db error", zap.Error(err))
		return
	}

	userRepository := user.NewUserRepository(db)
	userService := user.NewService(userRepository)

	apiCfg := httpapi.NewHttpAPI(&cfg.HttpConfig, userService)
	apiCfg.ListenAndServe()
}
