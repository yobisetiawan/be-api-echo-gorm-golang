package validator

import (
	"be_api/app/database"
	"be_api/app/models"
	"be_api/app/requests"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UniqueProductCategoryTitle(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	req, ok := fl.Parent().Interface().(requests.ProductCategoryRequest)
	if !ok {
		// Handle error if struct type doesn't match
		return false
	}

	// Check if exists in the database
	var data models.ProductCategory
	query := database.DB.Model(data).Where("title = ?", value)
	if req.ID != "" {
		query.Where("id != ?", req.ID)
	}

	result := query.First(&data)
	return result.Error == gorm.ErrRecordNotFound
}
