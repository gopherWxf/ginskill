package Getter

import (
	"fmt"
	"ginskill/src/data/mappers"
	"ginskill/src/models/UserModel"
	"ginskill/src/result"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	GetUserList() []*UserModel.UserModelImpl
	GetUserByID(id int) *result.ErrorResult
}

type UserGetterImpl struct {
	userMapper *mappers.UserMapper
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{userMapper: &mappers.UserMapper{}}
}

func (u *UserGetterImpl) GetUserList() (users []*UserModel.UserModelImpl) {
	u.userMapper.GetUserList().Query().Find(users)
	return
}
func (u *UserGetterImpl) GetUserByID(id int) *result.ErrorResult {
	user := UserModel.New()
	db := u.userMapper.GetUserDetail(id).Query().Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("not found user id:%d", id))
	}
	return result.Result(user, nil)
}
