package user

// LoginForm is used when a user logs in.
type LoginForm struct {
	Username   string `binding:"required,min=2" json:"username" form:"username"`
	Password   string `binding:"required,min=4" json:"password" form:"password"`
	RedirectTo string `binding:"-" form:"redirectTo" json:"-"`
	RememberMe string `binding:"-" form:"remember_me" json:"-"`
}

// AccountReset 密码重置
type AccountReset struct {
	Password      string `binding:"required,min=4" json:"old" form:"old"`
	NewPassword   string `binding:"required,min=4,nefield=Password" json:"new" form:"new"`
	CheckPassword string `binding:"required,min=4,eqfield=NewPassword" json:"confirm" form:"confirm"`
}

// UserinfoFormType is used when updating a user.
type UserinfoFormType struct {
	Username        string `validate:"required" form:"username" json:"username" needed:"true" len_min:"3" len_max:"20"`
	Email           string `json:"email" form:"email"`
	Language        string `validate:"default=en-us" form:"language" json:"language"`
	CurrentPassword string `validate:"omitempty,min=6,max=72" form:"current_password" json:"current_password" omit:"true"`
	Password        string `validate:"omitempty,min=6,max=72" form:"password" json:"password" len_min:"6" len_max:"72" equalInput:"ConfirmPassword"`
	ConfirmPassword string `validate:"omitempty" form:"password_confirmation" json:"password_confirmation" omit:"true"`
	Status          int    `validate:"default=0" form:"status" json:"status"`
	Theme           string `form:"theme" json:"theme"`
	AnidexAPIToken  string `validate:"-" form:"anidex_api" json:"anidex_api"`
	NyaasiAPIToken  string `validate:"-" form:"nyaasi_api" json:"nyaasi_api"`
	TokyoTAPIToken  string `validate:"-" form:"tokyot_api" json:"tokyot_api"`
}
