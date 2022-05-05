package model

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/vndee/saasuac/utils"
)

type ReturnParams struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetReturnParams(status string, message string, data interface{}) ReturnParams {
	params := ReturnParams{status, message, data}
	return params
}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{(*User)(nil)}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		utils.Panic(err)
	}

	return nil
}

func DropTables(db *pg.DB) error {
	fmt.Println(db)
	models := []interface{}{(*User)(nil)}
	for _, model := range models {
		exists, err := db.Model(model).Exists()
		if exists == false {
			continue
		}

		//err = db.DropTable(model, &orm.DropTableOptions{
		//	IfExists: true,
		//})
		utils.Panic(err)
	}

	return nil
}
