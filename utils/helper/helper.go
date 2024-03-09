package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, statusCode int, msg interface{}) {
	c.JSON(statusCode, gin.H{"response": msg})
}

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}

	return string(b)
}