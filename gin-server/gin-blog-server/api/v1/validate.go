package v1

import (
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    errmsg.SUCCESS,
		"message": "validate continue",
	})
}
