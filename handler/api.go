package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vndee/saasuac/model"
)

func Healthcheck(ctx *fiber.Ctx) error {
	return ctx.JSON(model.ReturnParams{"success", "I am living!", nil})
}
