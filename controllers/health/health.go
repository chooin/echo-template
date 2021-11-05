package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context)  {
	c.String(http.StatusOK, "OK")
}
