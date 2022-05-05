package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vndee/saasuac/config"
	"github.com/vndee/saasuac/model"
	"github.com/vndee/saasuac/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	params := new(model.RegisterParams)
	if err := ctx.BodyParser(params); err != nil {
		log.Panic(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ReturnParams{
			"error",
			"Bad request",
			nil})
	}

	// check if this email has already been registered
	user := model.User{Email: params.Email,
		Name: params.FirstName + " " + params.LastName}

	b, err := model.ExistsUserByPrimaryKey(&config.PostgreSQLConnection, user)
	utils.Panic(err)

	if b == true {
		return ctx.Status(fiber.StatusConflict).JSON(model.ReturnParams{
			"error",
			"This email has been registered!",
			user.Email})
	}

	password := []byte(params.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ReturnParams{
			"error",
			"Internal Server Error",
			nil})
	}

	user.user_id = uuid.New().String()

	user.Password = string(hashedPassword)
	id, err := model.InsertUser(&config.PostgreSQLConnection, user)
	if err != nil {
		log.Panic(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ReturnParams{
			"err",
			"Bad request",
			nil})
	}
	fmt.Println(id)
	fmt.Println(user)
	// save new record of the registered user

	// return jwt token
	return ctx.Status(fiber.StatusOK).JSON(model.ReturnParams{"success", "ok", nil})
}
