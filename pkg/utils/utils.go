package utils

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context, x interface{}) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		return
	}
}
