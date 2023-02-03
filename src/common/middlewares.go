package common

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				c.JSON(400, gin.H{"message": e})
			}
		}()
		c.Next()
	}
}
