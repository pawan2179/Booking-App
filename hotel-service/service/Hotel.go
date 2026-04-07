package service

import (
	"hotel-service/dto"
	"hotel-service/model"
	"hotel-service/repository"
)

type HotelService interface {
	GetAllHotels() ([]*model.Hotel, error)
	GetHotelById(payload *dto.GetHotelByIdRequest) (*model.Hotel, error)
	CreateHotel(payload *dto.CreateHotelRequest) (*model.Hotel, error)
	UpdateHotel(payload *dto.UpdateHotelRequest) (*model.Hotel, error)
	DeleteHotel(payload *dto.DeleteHotelRequest) error
}

type HotelServiceImpl struct {
	repo repository.HotelRepository
}

func NewHotelService(r *repository.HotelRepository) HotelService {
	return &HotelServiceImpl{
		repo: *r,
	}
}

func (s *HotelServiceImpl) GetAllHotels() ([]*model.Hotel, error) {
	return s.repo.GetAllHotels()
}

func (s *HotelServiceImpl) GetHotelById(payload *dto.GetHotelByIdRequest) (*model.Hotel, error) {
	return s.repo.GetHotelById(payload.Id)
}

func (s *HotelServiceImpl) CreateHotel(payload *dto.CreateHotelRequest) (*model.Hotel, error) {
	hotel, err := s.repo.CreateHotel(payload.Name, payload.Description)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

func (s *HotelServiceImpl) UpdateHotel(payload *dto.UpdateHotelRequest) (*model.Hotel, error) {
	hotel, err := s.repo.UpdateHotel(payload.Name, payload.Description, payload.Id)
	return hotel, err
}

func (s *HotelServiceImpl) DeleteHotel(payload *dto.DeleteHotelRequest) error {
	err := s.repo.DeleteHotelById(payload.Id)
	return err
}
