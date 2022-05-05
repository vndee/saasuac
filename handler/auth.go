package handler

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vndee/saasuac/config"
	"github.com/vndee/saasuac/model"
	"github.com/vndee/saasuac/utils"
	"golang.org/x/crypto/bcrypt"
)

func JwtTokenGeneration(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	return t, err
}

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

	user.Id = uuid.New().String()

	user.Password = string(hashedPassword)

	// save new user record
	_, err = model.InsertUser(&config.PostgreSQLConnection, user)
	if err != nil {
		log.Panic(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ReturnParams{
			"err",
			"Bad request",
			nil})
	}

	log.Println("Succesfully created user with email: ", user.Email)

	// generate jwt token
	t, err := JwtTokenGeneration(user)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ReturnParams{
			"error",
			"Internal Server Error",
			err})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.ReturnParams{"success", "Succesfully registered!", t})
}

func Login(ctx *fiber.Ctx) error {
	params := new(model.LoginParams)
	if err := ctx.BodyParser(params); err != nil {
		log.Panic(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ReturnParams{
			"error",
			"Error on login request",
			nil})
	}

	user := model.User{Email: params.Email}
	exist, err := model.ExistsUserByPrimaryKey(&config.PostgreSQLConnection, user)
	if exist == false {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ReturnParams{"success", "Email not found", nil})
	}

	user, err = model.SelectUserByPrimaryKey(&config.PostgreSQLConnection, user)
	if err != nil {
		log.Panic(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ReturnParams{"error",
			"Internal Server Error",
			nil})
	}

	if !utils.CheckPasswordHash(params.Password, user.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ReturnParams{"success", "Invalid password", nil})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ReturnParams{
			"error",
			"Internal Server Error",
			nil})
	}

	t, err := JwtTokenGeneration(user)
	if err != nil {
		log.Panic(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ReturnParams{
			"error",
			"Internal Server Error",
			nil})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(model.ReturnParams{
		"success",
		"Succesfully login",
		t})
}
