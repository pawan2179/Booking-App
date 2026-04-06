package service

import "hotel-service/repository"

type HotelService interface {
}

type HotelServiceImpl struct {
	repo repository.HotelRepository
}
