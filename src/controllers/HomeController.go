package controllers

import "github.com/gofiber/fiber/v2"

func WelcomeHandler(ctx *fiber.Ctx) error {

	return ctx.SendString("Welcome to Lema Service")

}
