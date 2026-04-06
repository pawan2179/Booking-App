package dto

type CreateHotelRequest struct {
	Name        string
	Description string
}

type CreateHotelResponse struct {
	Id          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}
