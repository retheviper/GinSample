package main

import (
	"./application/constants"
	"./application/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RouteHome(r.Group(constants.ApiBasePath))
	r.Run()
}
