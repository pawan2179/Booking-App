package controller

import (
	"encoding/json"
	"fmt"
	"hotel-service/dto"
	"hotel-service/service"
	"hotel-service/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HotelControllerImpl struct {
	service service.HotelService
}

func NewHotelController(_service service.HotelService) HotelControllerImpl {
	return HotelControllerImpl{
		service: _service,
	}
}

func (c *HotelControllerImpl) GetAllHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := c.service.GetAllHotels()
	if err != nil {
		return
	}

	fmt.Println("Fetched hotels: ", hotels)
	utils.WriteJsonSuccessResponse(w, hotels)
}

func (c *HotelControllerImpl) GetHotelById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		fmt.Println("Invalid to parse request: ", err)
		utils.WriteJsonErrorResponse(w, fmt.Errorf("Failed to get ID from request path"))
		return
	}
	payload := &dto.GetHotelByIdRequest{
		Id: id,
	}
	hotel, err := c.service.GetHotelById(payload)
	fmt.Println("Hotel fetched: ", hotel)
	utils.WriteJsonSuccessResponse(w, hotel)
}

func (c *HotelControllerImpl) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreateHotelRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Invalid request: ", err)
		return
	}
	fmt.Println(payload)
	c.service.CreateHotel(&payload)
}

func (c *HotelControllerImpl) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	var payload dto.UpdateHotelRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Invalid request body: ", err)
		return
	}
	c.service.UpdateHotel(&payload)
}

func (c *HotelControllerImpl) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In delete controller")
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	if err != nil {
		fmt.Println("Failed to fetch ID from request: ", err)
		return
	}
	payload := dto.DeleteHotelRequest{
		Id: id,
	}
	c.service.DeleteHotel(&payload)
}
