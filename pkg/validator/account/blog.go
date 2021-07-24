package account

// BlogForm 密码重置
type BlogForm struct {
	BlogName  string `binding:"required" json:"blogName" form:"blogName"`
	SubTitle  string `binding:"required" json:"subTitle" form:"subTitle"`
	BlogPhoto string `binding:"required" json:"blogProfilePhoto" form:"blogProfilePhoto"`
	BeiAn     string `json:"beiAn" form:"beiAn"`
	BTitle    string `json:"bTitle" form:"bTitle"`
}

type AccountForm struct {
	//  昵称
	Nickname string `json:"nickname,omitempty" form:"nickname"`
	// 头像 保存为链接
	Avatar string `binding:"required" json:"avatar" form:"avatar"`
	// 账户
	Email string `json:"email,omitempty" form:"email"`
	// 手机号
	Phone string `json:"phone,omitempty" form:"phone"`
	// 住址
	Address string `json:"address,omitempty" form:"address"`
	// 个人说明
	Profile string `json:"profile,omitempty" form:"profile"`
}
