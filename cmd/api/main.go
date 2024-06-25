package main

import (
	"international-trading-test/internal/app"
	"international-trading-test/internal/config"
)

func main() {
	cfg := config.GetConfig()
	mainApp := app.NewApp(cfg)
	mainApp.Start()
}
