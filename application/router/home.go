package router

import (
	"../constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteHome(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group(constants.ApiVersion)
	group.GET(constants.HomeUrl, homeContent)
}

func homeContent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "home",
	})
}
