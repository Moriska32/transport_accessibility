package rest

import (
	"log"
	"transport_accessibility_bus/configuration"
	"transport_accessibility_bus/datastore"
	"transport_accessibility_bus/env"
	"transport_accessibility_bus/rest/auth"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// MainAPI Главное API
func MainAPI(app *fiber.App, enviroment *env.Env, cfg *configuration.AppConfiguration) {
	/* Connection to database */
	dbConn := datastore.DataBase{DB: enviroment.DBs["main_db"]}

	/* Casbin */
	enforcerServices, err := casbin.NewEnforcer(cfg.RBACFileName)
	if err != nil {
		log.Fatalln("Casbin initialization error:", err)
	}
	err = InitCasbinPoliciesServices(enforcerServices, &dbConn)
	if err != nil {
		log.Fatalln("Can't query casbin policies from database;", err)
	}
	_ = enforcerServices

	/* JWT */
	jwtMVPBus := auth.InitAuth(&dbConn, cfg.JWTConfiguration)

	// API
	mainAPI := app.Group("/" + cfg.ServerCfg.APIPath)

	{
		mainAPI.Get("/test", MakeTest())
		mainAPI.Post("/doauth", jwtMVPBus.LoginHandler)
		mainAPI.Get("/refresh_token", jwtMVPBus.RefreshHandler)

		mainAPI.Get("/whoami", jwtMVPBus.MiddlewareFunc(), auth.GetMyRole(enviroment))
	}

}

//MakeTest Тестовая шляпа
func MakeTest() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("Hello, world!")
	}
}
