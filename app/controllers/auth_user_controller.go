package controllers

import (
	"be_api/app/database"
	"be_api/app/models"
	"be_api/app/requests"
	"be_api/app/utils"
	"be_api/app/validator"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type AuthUserController struct {
}

func NewAuthUserController() *AuthUserController {
	return &AuthUserController{}
}

func (ctr *AuthUserController) Register(c echo.Context) error {

	formData := requests.AuthRegisterRequest{}

	if err := c.Bind(&formData); err != nil {
		return utils.ErrorGeneralLog(c, 400, "Invalid Form Data", err)
	}

	if err := c.Validate(formData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validator.NewValidationErrorResponse(err, formData))
	}

	data := models.User{}
	if err := copier.Copy(&data, formData); err != nil {
		return utils.Error500Log(c, err)
	}

	hashPass, err := utils.HashPassword(data.Password)
	if err != nil {
		return utils.Error500Log(c, err)
	}

	data.Password = hashPass

	result := database.DB.Create(&data)
	if result.Error != nil {
		return utils.Error500Log(c, result.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": data})
}

func (ctr *AuthUserController) Login(c echo.Context) error {
	formData := requests.AuthLoginRequest{}

	if err := c.Bind(&formData); err != nil {
		return utils.ErrorGeneralLog(c, 400, "Invalid Form Data", err)
	}

	if err := c.Validate(formData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validator.NewValidationErrorResponse(err, formData))
	}

	dt := models.User{}

	result := database.DB.Where("email = ?", formData.Email).First(&dt)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Credential"})
	}

	if utils.HashCheckPassword(dt.Password, formData.Password) {

		accessToken, err := utils.JWTGenerateToken(jwt.MapClaims{
			"user_id": dt.ID,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		}, "")
		if err != nil {
			return err
		}

		session := models.Session{}
		database.DB.Where("user_id = ? and expired_at > ?", dt.ID, time.Now()).First(&session)

		refreshToken := session.RefreshToken

		if session.RefreshToken == "" {
			exp := time.Now().Add(time.Hour * 24 * 365)
			refreshToken, err = utils.JWTGenerateToken(jwt.MapClaims{
				"user_id": dt.ID,
				"exp":     exp.Unix(),
			}, "refresh")
			if err != nil {
				return err
			}
			sessiondt := models.Session{
				UserID:       dt.ID,
				IsActive:     true,
				ExpiredAt:    &exp,
				RefreshToken: refreshToken,
			}

			result := database.DB.Save(&sessiondt)
			if result.Error != nil {
				return utils.Error500Log(c, result.Error)
			}
		}

		markForDeleted := dt.MarkForDeletedAt
		if markForDeleted != nil && markForDeleted.After(time.Now()) {
			dt.MarkForDeletedAt = nil
			database.DB.Save(&dt)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Credential"})
	}
}
