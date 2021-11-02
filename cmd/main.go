package main

import (
	todoRouter "github.com/boomauakim/go-todo-clean-arch/todo/delivery/http/route"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	appEnv := viper.GetString("app.env")
	port := viper.GetString("app.port")

	todoRouter.SetupRouter(app)

	log.Fatal(app.Listen(":" + port))
}
