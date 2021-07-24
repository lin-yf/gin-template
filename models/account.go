package models

import (
	"errors"
	"time"

	"go-template/pkg/log"
	"go-template/pkg/utils"
	"go-template/setting"

	"gorm.io/gorm"
)

// Account 账户
type Account struct {
	gorm.Model
	// 账户名
	Username string `gorm:"column:username;uniqueIndex" json:"username"`
	Nickname string `json:"nickname"`
	// 账户密码
	Password string `json:"-"`
	// 二次验证token
	Token string `json:"-"`
	// 头像 保存为链接
	Avatar string `json:"avatar"`
	// 账户
	Email string `json:"email"`
	// 手机号
	Phone string `json:"phone"`
	// 住址
	Address string `json:"address"`
	// 个人说明
	Profile string `json:"profile"`
	// 创建时间
	CreateTime time.Time `json:"created"`
	// 最后登录时间
	LoginTime time.Time
	// 登出时间
	LogoutTime time.Time
	// 最后登录ip
	LoginIP string
}

// Blogger 博客信息
type Blogger struct {
	gorm.Model
	// 博客名
	BlogName string `json:"blogName"`
	// SubTitle
	SubTitle string `json:"subtitle"`
	// 头像 保存为链接
	BlogPhoto string
	// 备案号
	BeiAn string
	// 底部title
	BTitle string
	// 版权声明
	Copyright string
	// 主题样式
	Theme string
	// 主题样式路径
	ThemePath string
}

// LoadAccount 初始化加载账户信息
func LoadAccount() (*Account, error) {
	db := GetDB()
	OwlAccount := &Account{}
	err := db.Where("username=?", setting.Conf.Account.Username).First(OwlAccount).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		passwords, err := utils.SetPassword(setting.Conf.Account.Password)
		if err != nil {
			return OwlAccount, err
		}
		OwlAccount = &Account{
			Username:   setting.Conf.Account.Username,
			Password:   passwords,
			Email:      setting.Conf.Account.Email,
			Phone:      setting.Conf.Account.Phone,
			Address:    setting.Conf.Account.Address,
			CreateTime: time.Now(),
		}
		err = db.Create(OwlAccount).Error
		if err != nil {
			log.Errorf("初始化账号失败", err)
			return OwlAccount, err
		}
		return OwlAccount, nil
	}
	return OwlAccount, nil
}

// LoadBlogger 初始化加载博客信息
func LoadBlogger() (*Blogger, error) {
	db := GetDB()
	GlobalBlogInfo := &Blogger{}
	err := db.First(GlobalBlogInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		GlobalBlogInfo = &Blogger{
			BlogName:  "L的博客小站",
			SubTitle:  "coding and changing",
			Theme:     "owl",
			ThemePath: "usr/themes/",
		}
		err := db.Create(GlobalBlogInfo).Error
		if err != nil {
			log.Errorf("创建初始化博客信息失败:", err)
			return GlobalBlogInfo, err
		}
		return GlobalBlogInfo, nil
	}
	return GlobalBlogInfo, nil
}

// FindUserByName 查找用户
func (u *Account) FindUserByName(name string) error {
	db := GetDB()
	err := db.Where("username = ?", name).First(u).Error
	return err
}

// SetPassword 账户设置密码
func (u *Account) SetPassword(new string) {
	u.Password = new
}

// Update 更新账号信息
func (u *Account) Update(data map[string]interface{}) error {
	db := GetDB()
	err := db.Model(u).Updates(data).Error
	return err
}

// update 更新博客
func (bg *Blogger) Update(data map[string]interface{}) error {
	db := GetDB()
	err := db.Model(bg).Updates(data).Error
	return err
}
