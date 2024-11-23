package utils

import (
	"log"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
)

func CustomLog(v ...any) {
	// Get the caller's file and line number
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Printf("%s:%d: %s\n", file, line, v)
	} else {
		log.Println(v...)
	}
}

func Error500Log(c echo.Context, v ...any) error {
	// Get the caller's file and line number
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Printf("%s:%d: %s\n", file, line, v)
	} else {
		log.Println(v...)
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Something went wrong!"})
}

func ErrorGeneralLog(c echo.Context, code int, msg string, v ...any) error {
	// Get the caller's file and line number
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Printf("%s:%d: %s\n", file, line, v)
	} else {
		log.Println(v...)
	}

	return c.JSON(code, echo.Map{"message": msg})
}
