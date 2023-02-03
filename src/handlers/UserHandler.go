package handlers

import (
	"ginskill/src/models/UserModel"
	"ginskill/src/result"
	"ginskill/src/test"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := UserModel.New()
	result.Result(c.ShouldBind(user)).UnWrap()
	//{message:"xxx",code:"0001",result:xxx}
	OK(c)("userlist", "0000", result.Result(test.GetInfo(user.UserId)).UnWrap())
}
