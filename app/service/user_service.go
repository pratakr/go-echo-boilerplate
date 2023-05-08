package service

import "kancha-api/app/model"

func (s *Service) CreateUser(user *model.User) (*model.User, error) {
	tx := s.Db.Model(&model.User{})
	err := tx.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UpdateUser(user *model.User) (*model.User, error) {
	tx := s.Db.Model(&model.User{})
	err := tx.Debug().Where("id=?", user.ID).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) DeleteUser(id int64) error {
	tx := s.Db.Model(&model.User{})
	err := tx.Where("id=?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ExistUser(user *model.User) bool {
	tx := s.Db.Model(&model.User{})
	var count int64
	tx.Where("email=?", user.Email).Count(&count)
	return count > 0
}

func (s *Service) FindUsers() ([]*model.User, error) {
	var users []*model.User
	tx := s.Db.Model(&model.User{})
	err := tx.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) FindUserByID(id int64) (*model.User, error) {
	var user model.User
	tx := s.Db.Model(&model.User{})
	err := tx.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
