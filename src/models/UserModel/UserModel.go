package UserModel

type UserModelImpl struct {
	UserId   int    `json:"id"  form:"id"`
	UserName string `json:"name" form:"name" binding:"required,UserName"`
}

func (*UserModelImpl) TableName() string {
	return "users"
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
