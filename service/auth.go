package service

import (
	"fmt"

	"github.com/TravisRoad/gomarkit/global"
	"github.com/TravisRoad/gomarkit/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (as *AuthService) Login(username, password string) (model.User, error) {
	var user model.User
	if err := global.DB.Model(&model.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return user, fmt.Errorf("failed to find user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		user.Password = ""
		return user, fmt.Errorf("incorrect password: %w", err)
	}

	user.Password = ""
	return user, nil
}
