package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lema/src/actions"
)

func GetUsers(ctx *fiber.Ctx) error {
	return actions.GetUserHandler(ctx)
}

func GetUserCount(ctx *fiber.Ctx) error {
	return actions.GetUserCountHandler(ctx)
}

func GetUserById(ctx *fiber.Ctx) error {
	return actions.GetUserByIdHandler(ctx)
}
