package service

import (
	"errors"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/model"
	"gorm.io/gorm"
)

type ProjService struct{}

func (ps *ProjService) GetProj(page, size int) (projs []model.Proj, cnt int64, err error) {
	if err := global.DB.Model(&model.Proj{}).Count(&cnt).Error; err != nil {
		return nil, 0, err
	}

	if err := global.DB.Model(&model.Proj{}).Offset((page - 1) * size).Limit(size).Find(&projs).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, err
		}
	}

	return
}
