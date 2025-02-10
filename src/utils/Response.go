package utils

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type PaginateData struct {
	Data        interface{} `json:"data"`
	Total       int64       `json:"total"`
	Page        int         `json:"page"`
	PerPage     int         `json:"perPage"`
	TotalPages  int         `json:"totalPages"`
	HasPrevPage bool        `json:"hasPrevPage"`
	HasNextPage bool        `json:"hasNextPage"`
	PrevPage    int         `json:"prevPage"`
	NextPage    int         `json:"nextPage"`
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}, messages ...string) error {
	//ctx := new(fiber.Ctx)
	respMessage := "Success"
	if len(messages) > 0 {
		for _, msg := range messages {
			respMessage = respMessage + msg
		}
		respMessage = strings.Replace(respMessage, "Success", "", 1)
	}
	resp := make(map[string]interface{})
	resp["status"] = true
	resp["message"] = respMessage
	resp["code"] = 200
	resp["data"] = data
	resp["errors"] = nil
	return ctx.Status(200).JSON(resp)
}

func Ok(ctx *fiber.Ctx) error {
	resp := make(map[string]interface{})
	resp["status"] = true
	resp["message"] = "Success"
	return ctx.Status(200).JSON(resp)
}

func ErrorResponse(ctx *fiber.Ctx, data interface{}) error {
	resp := make(map[string]interface{})
	resp["status"] = false
	resp["message"] = data
	resp["code"] = 400
	resp["data"] = nil
	resp["errors"] = data
	return ctx.Status(400).JSON(resp)
}

func BadRequest(ctx *fiber.Ctx, data interface{}, messages ...string) error {

	respMessage := "Failed"
	if len(messages) > 0 {
		for _, msg := range messages {
			respMessage = respMessage + msg
		}
		respMessage = strings.Replace(respMessage, "Failed", "", 1)
	}

	if _, ok := data.(string); ok {
		respMessage = data.(string)
	}

	resp := make(map[string]interface{})
	resp["status"] = false
	resp["message"] = respMessage
	resp["code"] = 400
	resp["data"] = nil
	resp["errors"] = data
	return ctx.Status(400).JSON(resp)
}

func FailedResponse(ctx *fiber.Ctx, data interface{}, messages ...string) error {

	respMessage := "Failed"
	if len(messages) > 0 {
		for _, msg := range messages {
			respMessage = respMessage + msg
		}
		respMessage = strings.Replace(respMessage, "Failed", "", 1)
	}

	if _, ok := data.(string); ok {
		respMessage = data.(string)
	}

	resp := make(map[string]interface{})
	resp["status"] = false
	resp["message"] = respMessage
	resp["code"] = 407
	resp["data"] = nil
	resp["errors"] = data
	return ctx.Status(407).JSON(resp)
}
