package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/qinains/fastergoding"
	"github.com/vndee/saasuac/config"
	"github.com/vndee/saasuac/model"
	"github.com/vndee/saasuac/router"
)

func main() {
	commandPtr := flag.String("c", "serve", "Command need to be executed")
	flag.Parse()

	fmt.Println("Command:", *commandPtr)

	if *commandPtr == "create_database" {
		err := model.CreateSchema(&config.PostgreSQLConnection)
		if err != nil {
			fmt.Println("[FAILED] Failed to create database schema!")
			panic(err)
		} else {
			fmt.Println("[SUCCESS] Created database schema!")
		}
	} else if *commandPtr == "clear_database" {
		fmt.Println("clear")
		// TODO
		//err := model.DropTables(&config.PostgreSQLConnection)
		//if err != nil {
		//	fmt.Println("[FAILED] Failed to drop tables!")
		//	panic(err)
		//} else {
		//	fmt.Println("[SUCCESS] Dropped tables")
		//}
	} else if *commandPtr == "reset_database" {
		// TODO
		fmt.Println("reset")
	} else {
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

		router.SetupRoutes(app)

		bindingPort := fmt.Sprintf(":%s", config.Config("PORT"))
		log.Fatal(app.Listen(bindingPort))
	}
}
