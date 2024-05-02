package v1

import (
	"NyaLog/gin-blog-server/service"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 生成RSS内容
func GenerateRSS(c *gin.Context) {
	rssinfo := service.GenerateRSS()
	xmlBytes, _ := xml.Marshal(rssinfo)
	c.Data(http.StatusOK, "application/xml", xmlBytes)
}
