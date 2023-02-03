package UserModel

type UserModelImpl struct {
	UserId   int    `json:"id"`
	UserName string `json:"name" form:"name" binding:"required,min=4"`
}

func New(attrs ...UserModelAttrFunc) *UserModelImpl {
	u := &UserModelImpl{}
	UserModelAttrFuncs(attrs).Apply(u)
	return u
}
func (this *UserModelImpl) Mutate(attrs ...UserModelAttrFunc) *UserModelImpl {
	UserModelAttrFuncs(attrs).Apply(this)
	return this
}