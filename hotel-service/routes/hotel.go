package routes

import (
	"hotel-service/controller"

	"github.com/go-chi/chi/v5"
)

type HotelRouter struct {
	controller *controller.HotelControllerImpl
}

func NewHotelRouter(uc *controller.HotelControllerImpl) HotelRouter {
	return HotelRouter{
		controller: uc,
	}
}

func (r *HotelRouter) Register(router chi.Router) {
	router.Get("/", r.controller.GetAllHotels)
	router.Get("/{id}", r.controller.GetHotelById)
	router.Post("/create", r.controller.CreateHotel)
	router.Post("/update", r.controller.UpdateHotel)
	router.Delete("/", r.controller.DeleteHotel)
}
