package test

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetInfo(id int) (gin.H, error) {
	if id > 10 {
		return gin.H{"message": "test"}, nil
	} else {
		return nil, errors.New("test error")
	}
}
