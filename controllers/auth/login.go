package auth

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Login(c *gin.Context)  {
	 c.JSON(200, gin.H{
		"message": "OK",
		"gin_mode": os.Getenv("GIN_MODE"),
	})
}
