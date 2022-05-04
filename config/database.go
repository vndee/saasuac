package config

import (
	"github.com/go-pg/pg/v10"
)

var PostgreSQLConnection = *pg.Connect(&pg.Options{
	Addr:     Config("POSTGRES_ADDR"),
	User:     Config("POSTGRES_USER"),
	Password: Config("POSTGRES_PASSWORD"),
	Database: Config("POSTGRES_DB")})
