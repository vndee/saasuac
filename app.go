package main

import (
	"fmt"
	"log"
	"config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run()

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Ho_Chi_Minh",
	}))
	app.Use(cors.New())
	app.Get("/dashboard", monitor.New())

	fmt.Println("Hello, World")		
	log.Fatal(app.Listen(fmt.Sprint(":%s", config.Config("PORT"))))
}
