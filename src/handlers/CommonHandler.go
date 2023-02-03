package handlers

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type JSONResult struct {
	Message string      `json:"message,omitempty"`
	Code    string      `json:"code,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func NewJSONResult(message string, code string, result interface{}) *JSONResult {
	return &JSONResult{Message: message, Code: code, Result: result}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJSONResult("", "", nil)
		},
	}
}

type ResultFunc func(message string, code string, result interface{})

func OK(c *gin.Context) ResultFunc {
	return func(message string, code string, result interface{}) {
		r := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(r)
		r.Message = message
		r.Code = code
		r.Result = result
		c.JSON(200, r)
	}
}
