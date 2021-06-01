package main

import (
	"./application/constants"
	"./application/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RouteMember(r.Group(constants.ApiBasePath))
	err := r.Run()
	if err != nil {
		return
	}
}
