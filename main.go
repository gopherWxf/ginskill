package main

import (
	"ginskill/src/common"
	"ginskill/src/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(common.ErrorHandler())
	r.POST("/users", handlers.UserList)
	r.POST("/users:id", handlers.UserList)

	r.Run(":80")
}
