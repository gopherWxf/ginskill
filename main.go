package main

import (
	"ginskill/src/common"
	"ginskill/src/handlers"
	_ "ginskill/src/validators"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(common.ErrorHandler())
	r.GET("/users", handlers.UserList)
	r.GET("/users:id", handlers.UserDetail)
	r.POST("/users", handlers.UserSave)
	r.Run(":80")
}
