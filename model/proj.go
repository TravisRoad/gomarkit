package model

import (
	"gorm.io/gorm"
)

type Proj struct {
	gorm.Model
	UserID uint   `gorm:"comment:用户ID" json:"user_id"`
	Name   string `gorm:"type:varchar(32);comment:项目名" json:"name"`
	Desc   string `gorm:"type:varchar(128);comment:项目描述" json:"desc"`
	Type   string `gorm:"type:varchar(64);comment:项目类型，0为序列标注，1为其他" json:"type"`
	Detail string `gorm:"type:json;comment:项目详情" json:"detail"`
}

type ProjTag struct {
	gorm.Model
	ProjID uint   `gorm:"comment:项目ID" json:"proj_id"`
	Name   string `gorm:"type:varchar(32);comment:标签名" json:"name"`
}
