package model

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type User struct {
	tableName struct{} `pg:"Users"`
	Id        string   `pg:"user_id"`
	Name      string   `pg:"user_name"`
	Email     string   `pg:"email,pk"`
	Password  string   `pg:"password"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s %s>", u.Id, u.Name, u.Email, u.Password)
}

func ExistsUserByPrimaryKey(pg *pg.DB, user User) (bool, error) {
	id, err := pg.Model(&user).WherePK().Exists()
	return id, err
}

func InsertUser(pg *pg.DB, user User) (orm.Result, error) {
	id, err := pg.Model(&user).Insert()
	return id, err
}

func InsertUserBatch(pg *pg.DB, users []User) (orm.Result, error) {
	id, err := pg.Model(&users).Insert()
	return id, err
}

func SelectUserByPrimaryKey(pg *pg.DB, user User) (User, error) {
	err := pg.Model(&user).WherePK().Select()
	return user, err
}

func UpdateUserByPrimaryKey(pg *pg.DB, user User) (orm.Result, error) {
	id, err := pg.Model(&user).WherePK().Update()
	return id, err
}

func DeleteUserByPrimaryKey(pg *pg.DB, user User) (orm.Result, error) {
	id, err := pg.Model(&user).WherePK().Delete()
	return id, err
}
