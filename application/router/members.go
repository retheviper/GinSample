package router

import (
	"../constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteMember(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group(constants.MembersUrl)
	group.GET(constants.HomeUrl, listMember)
	group.GET(constants.HomeUrl+"/:id", getMember)
}

func listMember(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "listMember",
	})
}

func getMember(ctx *gin.Context) {
	//id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "getMember",
	})
}
