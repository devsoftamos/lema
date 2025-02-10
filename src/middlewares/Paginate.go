package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

var Page string
var PerPage string

func Paginate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Page = ctx.Query("pageNumber", "1")
		PerPage = ctx.Query("pageSize", "10")
		return ctx.Next()
	}
}
