package Setter

import (
	"ginskill/src/data/mappers"
	"ginskill/src/models/LogModel"
	"ginskill/src/models/UserModel"
	"ginskill/src/result"
	"time"
)

var UserSetter IUserSetter

func init() {
	UserSetter = NewUserSetterImpl()
}

type IUserSetter interface {
	SaveUser(user *UserModel.UserModelImpl) *result.ErrorResult
}

type UserSetterImpl struct {
	userMapper *mappers.UserMapper
	logMapper  *mappers.LogMapper
}

func NewUserSetterImpl() *UserSetterImpl {
	return &UserSetterImpl{userMapper: &mappers.UserMapper{}, logMapper: &mappers.LogMapper{}}
}
func (this *UserSetterImpl) SaveUser(user *UserModel.UserModelImpl) *result.ErrorResult {
	addUser := this.userMapper.AddNewUser(user)
	addLog := this.logMapper.AddLog(LogModel.NewLogImpl("add log", time.Now()))
	err := mappers.Mappers(addUser, addLog).Exec(func() error {
		err := addUser.Exec().Error
		if err != nil {
			return err
		}
		err = addLog.Exec().Error
		if err != nil {
			return err
		}
		return nil
	})

	return result.Result("success", err)
}
