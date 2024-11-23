package controllers

import (
	"be_api/app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserProfileController struct {
}

func NewUserProfileController() *UserProfileController {
	return &UserProfileController{}
}

// @Tags user-profile
// @Router /v1/user/profile [get]
// @Security BearerAuth
func (ctr *UserProfileController) Profile(c echo.Context) error {
	currentUser, err := utils.CurrentUser(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"data": &currentUser})
}
