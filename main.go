package main

import (
	"ginskill/src/common"
	"ginskill/src/models/UserModel"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(common.ErrorHandler())
	r.POST("/", func(c *gin.Context) {
		user := UserModel.New()
		if err := c.ShouldBind(user); err != nil {
			panic(err.Error())
		} else {
			c.JSON(200, user)
		}
	})
	r.Run(":80")
}
