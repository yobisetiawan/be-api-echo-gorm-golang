package appMiddleware

import "github.com/labstack/echo/v4"

func SetIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Assume the ID comes from the URL parameter
		id := c.Param("id")
		c.Set("ID", id) // Store it in the context
		return next(c)
	}
}
