package user

import (
	"larago/app/models"
	"larago/pkg/password"
)

// User 用户模型
type User struct {
	models.BaseModelId
	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password        string `gorm:"type:varchar(255)" valid:"password"`
	PasswordComfirm string `gorm:"-" valid:"password_comfirm"`
	models.BaseModelTime
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}
