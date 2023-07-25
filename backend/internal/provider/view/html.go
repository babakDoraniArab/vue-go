package view

import (
	"os"

	"github.com/gin-gonic/gin"
)

func WithGlobalData(data gin.H) gin.H {
	data["APP_NAME"] = os.Getenv("APP_NAME")
	return data
}
