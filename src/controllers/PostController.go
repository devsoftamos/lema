package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lema/src/actions"
	"lema/src/requests"
	"lema/src/utils"
)

func CreatePost(ctx *fiber.Ctx) error {

	createPostRequest := requests.CreatePostRequest{}
	createPost, err := createPostRequest.FromJSON(ctx.Body())

	if err != nil {
		return utils.BadRequest(ctx, err.Error())
	}
	err = createPost.Validate()
	if err != nil {
		return utils.BadRequest(ctx, err.Error())
	}
	return actions.CreatePostHandler(ctx, createPost)
}

func GetPosts(ctx *fiber.Ctx) error {
	return actions.GetPostsHandler(ctx)
}

func DeletePost(ctx *fiber.Ctx) error {
	return actions.DeletePostsHandler(ctx)
}
