package model

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type ReturnParams struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    map[string]*string `json:"data"`
}

func GetReturnParams(status string, message string, data map[string]*string) ReturnParams {
	params := ReturnParams{status, message, data}
	return params
}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{(*User)(nil)}
	for _, model := range models {
		exists, err := db.Model(model).Exists()
		fmt.Println(exists)
		fmt.Println(err)
		if exists == true {
			continue
		}

		err = db.Model(model).CreateTable(&orm.CreateTableOptions{})

		if err != nil {
			return err
		}
	}

	return nil
}
