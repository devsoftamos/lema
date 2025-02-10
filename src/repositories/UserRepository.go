package repositories

import (
	"lema/src/database"
	"lema/src/models"
)

type UserRepository struct {
}

func (user *UserRepository) GetPaginatedUsers(offset int, limit int) ([]models.User, int64, error) {
	var result []models.User
	query := database.SqliteDB.
		Offset(offset).
		Limit(limit).
		Model(&models.User{}).
		Joins("Address").
		Preload("Address").
		Order("users.id desc")
	query.Scan(&result)

	var count int64
	database.SqliteDB.Model(&models.User{}).
		Count(&count)

	return result, count, query.Error
}

func (user *UserRepository) GetUserCount() int64 {
	var count int64
	database.SqliteDB.Model(&models.User{}).Count(&count)
	return count
}

func (user *UserRepository) GetUserById(id uint) (models.User, error) {
	var result models.User
	query := database.SqliteDB.
		Model(&models.User{}).
		Joins("Address").
		Preload("Address").
		Where("users.id = ?", id)
	query.Scan(&result)
	return result, query.Error
}
