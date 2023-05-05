package service

import (
	"errors"
	"kancha-api/app/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) FindUserByEmail(email string, appId int) (*models.User, error) {
	var user *models.User
	tx := s.Db.Model(&models.User{})
	tx = tx.Where("email=?", email)
	if err := tx.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UserLogin(email string, password string) (*models.User, error) {
	var user *models.User
	tx := s.Db.Model(&models.User{})
	tx = tx.Where("email=?", email)
	err := tx.First(&user).Error
	if err != nil {
		return nil, err
	}
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("password not match")
	}

	return user, nil
}
