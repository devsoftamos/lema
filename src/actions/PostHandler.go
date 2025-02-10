package actions

import (
	"github.com/gofiber/fiber/v2"
	"lema/src/models"
	"lema/src/repositories"
	"lema/src/requests"
	"lema/src/utils"
)

var postRepository repositories.PostRepository

func CreatePostHandler(ctx *fiber.Ctx, createPostRequest requests.CreatePostRequest) error {

	post, _ := postRepository.CreatePost(models.Post{
		UserId: createPostRequest.UserId,
		Title:  createPostRequest.Title,
		Body:   createPostRequest.Body,
	})

	if post.ID <= 0 {
		return utils.BadRequest(ctx, "unable to create post")
	}
	return utils.SuccessResponse(ctx, post, "Post created")

}

func GetPostsHandler(ctx *fiber.Ctx) error {
	posts, _ := postRepository.GetPosts(ctx)
	return utils.SuccessResponse(ctx, posts)
}

func DeletePostsHandler(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.BadRequest(ctx, nil, "Invalid parameter")
	}

	post, _ := postRepository.GetPostById(uint(id))
	if post.ID <= 0 {
		return utils.BadRequest(ctx, nil, "Post not found")
	}

	postRepository.DeletePost(post)
	return utils.SuccessResponse(ctx, nil, "post deleted")

}
