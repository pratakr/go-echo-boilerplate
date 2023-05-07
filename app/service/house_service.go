package service

import "kancha-api/app/model"

func (s *Service) FindHouses() ([]*model.House, error) {
	var houses []*model.House
	tx := s.Db.Model(&model.House{})
	err := tx.Find(&houses).Error
	if err != nil {
		return nil, err
	}
	return houses, nil
}

func (s *Service) FindHouseById(id string) (*model.House, error) {
	var house *model.House
	tx := s.Db.Model(&model.House{})
	err := tx.Where("id=?", id).First(&house).Error
	if err != nil {
		return nil, err
	}
	return house, nil
}

func (s *Service) CreateHouse(house *model.House) error {
	tx := s.Db.Model(&model.House{})
	err := tx.Create(&house).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateHouse(house *model.House) error {
	tx := s.Db.Model(&model.House{})
	err := tx.Where("id=?", house.ID).Updates(&house).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteHouse(id int32) error {
	tx := s.Db.Model(&model.House{})
	err := tx.Where("id=?", id).Delete(&model.House{}).Error
	if err != nil {
		return err
	}
	return nil
}
