package appMiddleware

import (
	"be_api/app/configs"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JWT middleware function
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "missing auth token",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return []byte(configs.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid or expired token",
			})
		}

		if token == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token claims",
			})
		}

		userID, ok := claims["user_id"].(float64)
		if ok {
			c.Set("user_id", userID)
		}

		adminID, ok := claims["admin_id"].(float64)
		if ok {
			c.Set("admin_id", adminID)
		}

		if adminID == 0 && userID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized",
			})
		}

		return next(c)
	}
}
