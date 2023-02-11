package validators

import (
	"github.com/go-playground/validator/v10"
	"log"
)

func init() {
	if err := myvalid.RegisterValidation("UserName", VUserName); err != nil {
		log.Fatal("validator username err")
	}
}

var VUserName validator.Func = func(fl validator.FieldLevel) bool {
	uname, ok := fl.Field().Interface().(string)
	if ok && len(uname) >= 4 {
		return true
	}
	return false
}
