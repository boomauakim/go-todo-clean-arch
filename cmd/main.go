package main

import (
	"log"
	"strings"

	todoRouter "github.com/boomauakim/go-todo-clean-arch/todo/delivery/http/route"
	"github.com/boomauakim/go-todo-clean-arch/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Fatal error config file: %s \n", err)
		}
	}
	appEnv := viper.GetString("app.env")
	port := viper.GetString("app.port")

	var logger *zap.Logger
	var err error
	if appEnv == "production" {
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = config.Build()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			utils.ErrorHandler(c, e)
			return nil
		},
	})

	todoRouter.RegisterRouter(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"message": "The requested url doesn't exist.",
			},
		})
	})

	log.Fatal(app.Listen(":" + port))
}
