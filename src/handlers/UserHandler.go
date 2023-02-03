package handlers

import (
	"ginskill/src/models/UserModel"
	"ginskill/src/result"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := UserModel.New()
	result.Result(c.ShouldBind(user)).UnWrap()
	//{message:"xxx",code:"0001",result:xxx}
	if user.UserId > 10 {
		R(c)("userlist", "0000", "userlist")(Ok2String)
	} else {
		R(c)("userlist", "0000", "userlist")(Error)
	}
}
