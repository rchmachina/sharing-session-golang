package helper

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)


func JSONResponse(c echo.Context, statusCode int, msg interface{}) error {
	
	return c.JSON(statusCode, map[string]interface{}{
		"message": msg,
	})
}

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}

	return string(b)
}