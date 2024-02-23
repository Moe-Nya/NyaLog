package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": 111,
	})
}
