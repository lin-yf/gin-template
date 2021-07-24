package models

// User 用户表
type User struct {
	Username     string  `gorm:"column:username;uniqueIndex"`
	Email        string  `gorm:"column:email;uniqueIndex"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

// FindUserByName 查找用户
func (u *User) FindUserByName(name string) error {
	db := GetDB()
	err := db.Where("username = ?", name).First(u).Error
	return err
}
