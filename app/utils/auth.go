package utils

import (
	"be_api/app/database"
	"be_api/app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CurrentUser(c echo.Context) (models.User, error) {

	user_id := c.Get("user_id").(float64)
	if user_id == 0 {
		return models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	user := models.User{}
	result := database.DB.Model(models.User{}).Where("id = ?", user_id).First(&user)
	if result.Error != nil {
		return models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	return user, nil
}
