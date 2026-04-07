package dto

type GetHotelByIdRequest struct {
	Id int64 `json:"id"`
}

type CreateHotelRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateHotelResponse struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type UpdateHotelRequest struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteHotelRequest struct {
	Id int64 `json:"id"`
}
