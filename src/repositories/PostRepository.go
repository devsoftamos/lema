package repositories

import (
	"github.com/gofiber/fiber/v2"
	"lema/src/database"
	"lema/src/models"
)

type PostRepository struct {
}

func (post *PostRepository) CreatePost(pt models.Post) (models.Post, error) {
	var result models.Post
	query := database.SqliteDB.Model(&models.Post{}).
		Create(&pt)
	query.Scan(&result)
	return result, query.Error
}

func (post *PostRepository) GetPosts(ctx *fiber.Ctx) ([]models.Post, error) {
	userId := ctx.Query("userId")

	var result []models.Post
	query := database.SqliteDB.
		Model(&models.Post{})

	if userId != "" {
		query.Where("posts.user_id = ?", userId)
	}

	query.Scan(&result)
	return result, query.Error

}

func (post *PostRepository) GetPostById(id uint) (models.Post, error) {
	var result models.Post
	query := database.SqliteDB.
		Model(&models.Post{}).
		Where("posts.id = ?", id)
	query.Scan(&result)
	return result, query.Error
}

func (post *PostRepository) DeletePost(pt models.Post) {
	database.SqliteDB.Delete(&pt)
}
