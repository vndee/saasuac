package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vndee/saasuac/model"
)

func Register(ctx *fiber.Ctx) error {
	var params model.RegisterParams
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ReturnParams{
			"error",
			"Bad request",
			nil})
	}

	fmt.Println(params)
	// check if this email has already been registered

	// save new record of the registered user

	// return jwt token
	return ctx.Status(fiber.StatusOK).JSON(model.ReturnParams{"success", "ok", nil})
}
