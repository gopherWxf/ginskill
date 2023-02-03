package handlers

import (
	"ginskill/src/data/Getter"
	"ginskill/src/result"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	R(c)("userlist", "0000", Getter.UserGetter.GetUserList())(OK)

}
func UserDetail(c *gin.Context) {
	id := &struct {
		Id int `uri:"id" binding:"required,gt=0"`
	}{}
	result.Result(c.ShouldBindUri(id)).UnWrap()
	R(c)("userdetail", "0001", Getter.UserGetter.GetUserByID(id.Id).UnWrap())(OK)
}
