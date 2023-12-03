package service

import (
	"fmt"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// GetUsers retrieves a list of users from the database based on the specified pagination parameters.
//
// page: The page number of the results to retrieve. start from 1
// size: The number of results per page.
// []model.User: An array of model.User objects representing the retrieved users.
func (us *UserService) GetUsers(page, size int) ([]model.User, int64, error) {
	users := make([]model.User, 0)
	var cnt int64

	tx := global.DB.Begin()

	if err := tx.Model(&model.User{}).Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	if err := tx.Model(&model.User{}).Count(&cnt).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	for _, u := range users {
		u.Password = ""
	}

	tx.Commit()
	return users, cnt, nil
}

func (us *UserService) AddUser(u model.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt generate error: %w", err)
	}
	user := model.User{
		Username: u.Username,
		Password: string(hashedPass),
		Role:     global.ROLE_USER,
	}

	var cnt int64
	tx := global.DB.Begin()
	if err := tx.Model(&model.User{}).Where("username = ?", u.Username).Count(&cnt).Error; err != nil {
		tx.Rollback()
		return err
	}
	if cnt > 0 {
		return fmt.Errorf("user already exists")
	}
	if err := tx.Model(&model.User{}).Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (us *UserService) DeleteUser(id uint) error {
	if err := global.DB.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(u model.User) error {
	var user model.User
	tx := global.DB.Begin()

	if err := tx.Model(&model.User{}).Where("id = ?", u.ID).Take(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt generate error: %w", err)
	}

	// update
	user.Password = string(hashedPass)
	if len(u.Role) > 0 {
		user.Role = u.Role
	}

	if err := tx.Model(&model.User{}).Where("id = ?", u.ID).Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
