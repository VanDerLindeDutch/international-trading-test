package app

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "international-trading-test/docs"
	"international-trading-test/internal/config"
	"international-trading-test/internal/handler/http"
	"international-trading-test/internal/service"
	"international-trading-test/internal/service/redis-cache"
	"log"
)

type App struct {
	cfg    *config.Config
	fiber  *fiber.App
	logger *log.Logger
}

// @title Bis API
// @version 1.0

// @host 127.0.0.1:8081
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewApp(config *config.Config) App {

	app := fiber.New(fiber.Config{DisableStartupMessage: !config.IsDebug})

	app.Use(recover2.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
	}))

	return App{
		cfg:    config,
		fiber:  app,
		logger: log.Default(),
	}
}

func (app *App) Start() {
	ctx := context.Background()
	app.registerHandler(ctx)
	app.logger.Printf("Server start listening on port:%s", app.cfg.Listen.Port)

	err := app.fiber.Listen(fmt.Sprintf("%s:%s", app.cfg.Listen.BindIP, app.cfg.Listen.Port))
	if err != nil {
		panic(err)
	}

}

func (app *App) registerHandler(ctx context.Context) {
	log.Println(app.cfg)
	cache, err := redis_cache.New(ctx, app.cfg.Redis.Address, app.cfg.Redis.Password, 0)
	if err != nil {
		log.Fatalf("error while connecting to redis:%v", err)
	}
	svc := service.NewService(cache)
	handler := http_handler.NewHandler(app.cfg, svc)

	handler.Register(app.fiber)

}
