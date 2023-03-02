package main

import (
	"log"

	"github.com/srfbogomolov/warehouse_api/internal/config"
	"github.com/srfbogomolov/warehouse_api/internal/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	unsugared, err := logger.CreateLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = unsugared.Sugar()
}
