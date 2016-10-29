package web

import (
	"github.com/gin-gonic/gin"
	"nicosoft.org/wechat/web/controller"
)

type Router struct {
	Addr string
}

func initRouteList() *gin.Engine {

	router := gin.Default()

	router.GET("/core", controller.WeChatInit)
	router.POST("/core", controller.WeChatService)

	router.POST("/manager/menu", controller.WeChatMenuCreate)

	return router
}

func (r *Router) Start() {

	if r.Addr != "" {
		initRouteList().Run(r.Addr)
	} else {
		initRouteList().Run()
	}

}
