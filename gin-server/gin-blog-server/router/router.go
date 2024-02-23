package router

import (
	v1 "NyaLog/gin-blog-server/api/v1"
	"NyaLog/gin-blog-server/utils"

	"github.com/gin-gonic/gin"
)

func IniRouter() {
	gin.SetMode(utils.AppMode)
	routerv1 := gin.Default()

	r1 := routerv1.Group("api/v1")
	r1.GET("hello", v1.StartServer)

	_ = routerv1.Run(utils.HttpPort)
}
