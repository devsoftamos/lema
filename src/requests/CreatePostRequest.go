package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"lema/src/database"
	"lema/src/models"
)

type CreatePostRequest struct {
	Title  string `json:"title" validate:"required,min=3,max=100,unique"`
	Body   string `json:"body" validate:"required,min=3,max=400"`
	UserId uint   `json:"user_id" validate:"required,exists"`
}

func (r *CreatePostRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("unique", uniqueTitle)
	validate.RegisterValidation("exists", userIdExists)

	if err := validate.Struct(r); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errArray := []string{}
		for _, validationError := range validationErrors {
			if validationError.Field() == "UserId" && validationError.Tag() == "exists" {
				errArray = append(errArray, "User ID does not exist in the system.")
			} else if validationError.Field() == "Title" && validationError.Tag() == "unique" {
				errArray = append(errArray, "A post with this title already exists.")
			} else if validationError.Field() == "Title" && validationError.Tag() == "min" {
				errArray = append(errArray, "Title must be at least 3 characters long.")
			} else if validationError.Field() == "Title" && validationError.Tag() == "max" {
				errArray = append(errArray, "Title must be less than 100 characters long.")
			} else if validationError.Field() == "Body" && validationError.Tag() == "min" {
				errArray = append(errArray, "Body must be at least 3 characters long.")
			} else if validationError.Field() == "Body" && validationError.Tag() == "max" {
				errArray = append(errArray, "Body must be less than 400 characters long.")
			} else {
				msg := fmt.Sprintf("Field '%s' failed validation: %s", validationError.Field(), validationError.ActualTag())
				errArray = append(errArray, msg)
			}
		}
		return errors.New(fmt.Sprintf("Validation errors: %s", errArray))
	}
	return nil
}

func uniqueTitle(fl validator.FieldLevel) bool {
	var post models.Post
	title := fl.Field().String()
	result := database.SqliteDB.Model(&models.Post{}).Where("title = ?", title).First(&post)
	return result.RowsAffected == 0
}

func userIdExists(fl validator.FieldLevel) bool {
	var user models.User
	userId := fl.Field().Uint()
	result := database.SqliteDB.Model(&models.User{}).Where("id = ?", userId).First(&user)
	return result.RowsAffected > 0
}

func (r *CreatePostRequest) FromJSON(body []byte) (CreatePostRequest, error) {
	register := CreatePostRequest{}
	buf := bytes.NewBuffer(body)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&register)
	if err != nil {
		return CreatePostRequest{}, err
	}
	return register, nil
}
