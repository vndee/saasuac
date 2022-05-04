package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/qinains/fastergoding"
	"github.com/vndee/saasuac/config"
	"github.com/vndee/saasuac/router"
)

func main() {
	fastergoding.Run()

	//err := model.CreateSchema(&config.PostgreSQLConnection)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println("Created database schema")

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Ho_Chi_Minh",
	}))
	app.Use(cors.New())
	app.Get("/dashboard", monitor.New())

	router.SetupRoutes(app)

	bindingPort := fmt.Sprintf(":%s", config.Config("PORT"))
	log.Fatal(app.Listen(bindingPort))
}
