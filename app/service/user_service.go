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
	err := tx.Where("id=?", user.ID).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) DeleteUser(id int32) error {
	tx := s.Db.Model(&model.User{})
	err := tx.Where("id=?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
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
