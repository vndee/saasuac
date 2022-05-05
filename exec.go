package main

import (
	"flag"
	"fmt"

	"github.com/vndee/saasuac/config"
	"github.com/vndee/saasuac/model"
)

func main() {
	commandPtr := flag.String("c", "create_database", "Command need to be executed")
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
		fmt.Println("Unknown command")
	}
}
