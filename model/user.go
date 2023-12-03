package model

import "gorm.io/gorm"

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);comment:用户名" json:"username"`
	Password string `gorm:"type:varchar(128);comment:密码" json:"password"`
	Role     string `gorm:"type:varchar(32);comment:角色" json:"role"`
}
