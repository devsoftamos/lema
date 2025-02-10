package actions

import (
	"github.com/gofiber/fiber/v2"
	"lema/src/middlewares"
	"lema/src/models"
	"lema/src/repositories"
	"lema/src/utils"
	"math"
	"strconv"
)

var userRepository = repositories.UserRepository{}

func GetUserHandler(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(middlewares.Page)
	perPage, _ := strconv.Atoi(middlewares.PerPage)
	limit := perPage
	offset := (page - 1) * perPage

	users, count, _ := userRepository.GetPaginatedUsers(offset, limit)

	if len(users) == 0 {
		return utils.SuccessResponse(ctx, []models.User{}, "success")
	}

	hasNextPage := page < int(math.Ceil(float64(count)/float64(perPage)))
	nextPage := 0
	if hasNextPage {
		nextPage = page + 1
	}

	responseData := utils.PaginateData{
		Total:       count,
		Data:        users,
		Page:        page,
		PerPage:     perPage,
		TotalPages:  int(math.Ceil(float64(count) / float64(perPage))),
		HasPrevPage: page-1 > 0,
		HasNextPage: hasNextPage,
		PrevPage:    page - 1,
		NextPage:    nextPage,
	}

	return utils.SuccessResponse(ctx, responseData, "Success")

}

func GetUserCountHandler(ctx *fiber.Ctx) error {
	result := userRepository.GetUserCount()
	responseData := fiber.Map{
		"count": result,
	}
	return utils.SuccessResponse(ctx, responseData, "Success")

}

func GetUserByIdHandler(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.BadRequest(ctx, nil, "Invalid parameter")
	}

	user, _ := userRepository.GetUserById(uint(id))

	if user.ID <= 0 {
		return utils.BadRequest(ctx, nil, "Record not found")
	}

	return utils.SuccessResponse(ctx, user, "Success")

}
