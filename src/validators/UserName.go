package validators

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	registerValidation("UserName", UserNameRole("required,min=4").toFunc())
}

type UserNameRole string

func (this UserNameRole) toFunc() validator.Func {
	validatorError["UserName"] = "用户名必须在4-8位之间"
	return func(fl validator.FieldLevel) bool {
		v, ok := fl.Field().Interface().(string)
		if ok {
			return this.validate(v)
		}
		return false
	}
}
func (this UserNameRole) validate(v string) bool {
	//默认的验证
	if err := myvalid.Var(v, string(this)); err != nil {
		return false
	}
	//自己自定义
	if len(v) > 8 {
		return false
	}

	return true
}
