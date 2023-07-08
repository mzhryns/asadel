package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zhryn/asadel/app/usecase/url_shortener"
	"github.com/zhryn/asadel/config"
	"github.com/zhryn/asadel/handler"
	"github.com/zhryn/asadel/pkg/mongodb"
	"github.com/zhryn/asadel/repository/url_repository_v1"
)

func main() {
	cfg := config.New()

	if cfg.Environment == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	db, err := mongodb.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, context.TODO())
	if err != nil {
		log.Panicln("Failed to initialized MongoDB:", err)
	}

	url_repo := url_repository_v1.New(db, cfg.DBName, cfg.DBUrlCollection)
	url_shortener_uc := url_shortener.NewUseCase(url_repo)

	h := handler.NewHandler(url_shortener_uc)

	router.GET("/{shortenUrl}", h.Resolve)
	router.POST("/shorten", h.Shorten)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
