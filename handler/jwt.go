package handler

import (
	"time"

	"github.com/vndee/saasuac/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/vndee/saasuac/model"
)

// Protected protect routes
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(model.ReturnParams{
			"error",
			"Missing or malformed JWT",
			nil})

	} else {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(model.ReturnParams{
			"error",
			"Invalid or expired JWT",
			nil})
	}
}

func VerifyToken(ctx *fiber.Ctx) error {
	accessToken := string(ctx.Request().Header.Peek("Authorization"))[7:]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET")), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ReturnParams{
			"error",
			"Internel Servier Error",
			nil})
	}

	newToken := jwt.New(jwt.SigningMethodHS256)
	newClaims := newToken.Claims.(jwt.MapClaims)
	newClaims["email"] = claims["email"]
	newClaims["user_id"] = claims["user_id"]
	newClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := newToken.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return ctx.JSON(model.ReturnParams{
			"error",
			"Internal Server Error",
			nil})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.ReturnParams{
		"success",
		"Token is valid",
		map[string]interface{}{
			"token":   t,
			"user_id": claims["user_id"],
			"email":   claims["email"]}})
}
