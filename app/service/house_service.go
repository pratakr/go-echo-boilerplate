package service

import "kancha-api/app/models"

func (s *Service) FindHouses() ([]*models.House, error) {
	var houses []*models.House
	tx := s.Db.Model(&models.House{})
	err := tx.Find(&houses).Error
	if err != nil {
		return nil, err
	}
	return houses, nil
}

func (s *Service) FindHouseById(id string) (*models.House, error) {
	var house *models.House
	tx := s.Db.Model(&models.House{})
	err := tx.Where("id=?", id).First(&house).Error
	if err != nil {
		return nil, err
	}
	return house, nil
}
