package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"transport_accessibility_bus/configuration"
	"transport_accessibility_bus/env"
	"transport_accessibility_bus/rest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title API Портала транспортной доступности
// @version 0.1.0

// @contact.name Поддержка API
// @contact.url http://mgtniip.ru/
// @contact.email it@mgtniip.ru

// @host 0.0.0.0:7938
// @BasePath /

// @schemes http https

// @securityDefinitions.apikey AuthJWT
// @in query
// @name token

var (
	confName = flag.String("conf", "", "configuration_examples/localhost.json")
)

func main() {

	// flag.Parse()
	// if *confName == "" {
	// 	flag.Usage()
	// 	return
	// }

	appConfiguration, err := configuration.NewConfiguration("configuration_examples/localhost.json")
	if err != nil {
		fmt.Printf("Can't start server due the error: %s\n", err.Error())
		return
	}

	env := env.PrepareEnv(appConfiguration)

	config := fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Println(err)
			return ctx.Status(500).JSON(fmt.Errorf("panic error"))
		},
		IdleTimeout: 10 * time.Second,
	}

	server := fiber.New(config)
	if appConfiguration.UseCORS {
		allCors := cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowHeaders:     "Origin, Authorization, Content-Type, Content-Length, Accept, Accept-Encoding, X-HttpRequest",
			AllowMethods:     "GET, POST, PUT, DELETE",
			ExposeHeaders:    "Content-Length",
			AllowCredentials: true,
			MaxAge:           5600,
		})
		server.Use(allCors)
	}

	server.Static(fmt.Sprintf("/%s/docs", appConfiguration.ServerCfg.APIPath), appConfiguration.DocsFolder)

	rest.MainAPI(server, env, appConfiguration)

	err = server.Listen(fmt.Sprintf(":%s", appConfiguration.ServerCfg.Port))
	if err != nil {
		fmt.Printf("Can't start server due the error: %s\n", err.Error())
	}
}
