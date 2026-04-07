package repository

import (
	"database/sql"
	"fmt"
	"hotel-service/model"
)

type HotelRepository interface {
	GetAllHotels() ([]*model.Hotel, error)
	GetHotelById(id int64) (*model.Hotel, error)
	CreateHotel(name string, description string) (*model.Hotel, error)
	UpdateHotel(name string, description string, id int64) (*model.Hotel, error)
	DeleteHotelById(id int64) error
}

type HotelRepositoryImpl struct {
	db *sql.DB
}

func NewHotelRepositoryImpl(_db *sql.DB) HotelRepository {
	return &HotelRepositoryImpl{
		db: _db,
	}
}

func (r *HotelRepositoryImpl) GetAllHotels() ([]*model.Hotel, error) {
	query := "SELECT id, name, description FROM hotel;"

	row, err := r.db.Query(query)
	if err != nil {
		fmt.Println("Failed to Get all hotels: ", err)
		return nil, err
	}
	defer row.Close()
	hotels := []*model.Hotel{}

	for row.Next() {
		hotel := &model.Hotel{}
		if err := row.Scan(&hotel.Id, &hotel.Name, &hotel.Description); err != nil {
			fmt.Println("Error in parsing Hotel response: ", err)
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (r *HotelRepositoryImpl) GetHotelById(id int64) (*model.Hotel, error) {
	query := "SELECT id, name, description FROM hotel where id = ?;"
	row := r.db.QueryRow(query, id)

	if row == nil {
		fmt.Println("Could not find any hotel with given ID: ", id)
		return nil, fmt.Errorf("No Hotels found with given ID")
	}

	hotel := &model.Hotel{}
	scanErr := row.Scan(&hotel.Id, &hotel.Name, &hotel.Description)
	if scanErr != nil {
		fmt.Println("Error in mapping fetched hotel to hotel model: ", scanErr)
		return nil, scanErr
	}

	return hotel, nil
}

func (r *HotelRepositoryImpl) CreateHotel(name string, description string) (*model.Hotel, error) {
	query := `INSERT INTO HOTEL (name, description) VALUES (?, ?);`

	result, err := r.db.Exec(query, name, description)

	if err != nil {
		fmt.Println("Error in creating hotel: ", err)
		return nil, err
	}

	insertedId, insertErr := result.LastInsertId()
	if insertErr != nil {
		fmt.Println("Failed to insert new Hotel in DB :", err)
		return nil, insertErr
	}

	searchQuery := `SELECT id, name, description FROM HOTEL where id = ?;`
	row := r.db.QueryRow(searchQuery, insertedId)

	hotel := &model.Hotel{}
	scanErr := row.Scan(&hotel.Id, &hotel.Name, &hotel.Description)

	if scanErr != nil {
		fmt.Println("Error in getting Hotel from DB: ", scanErr)
		return nil, scanErr
	}

	return hotel, nil
}

func (r *HotelRepositoryImpl) UpdateHotel(name string, description string, id int64) (*model.Hotel, error) {
	query := `UPDATE HOTEL SET name = ?, description = ?, updated_at = NOW() WHERE id = ?`
	_, err := r.db.Exec(query, name, description, id)

	if err != nil {
		fmt.Println("Failed to update Hotel in DB: ", err)
		return nil, err
	}

	return &model.Hotel{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   "",
		UpdatedAt:   "",
	}, nil
}

func (r *HotelRepositoryImpl) DeleteHotelById(id int64) error {
	query := `DELETE FROM HOTEL WHERE id = ?`
	row, err := r.db.Exec(query, id)

	if err != nil {
		fmt.Println("Cannot find hotel with this id in DB: ", err)
		return err
	}

	rowAffected, rowAffErr := row.RowsAffected()
	if rowAffErr != nil {
		fmt.Println("Failed to DELETE Hotel from DB: ", rowAffErr)
		return rowAffErr
	}

	if rowAffected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("Deleted Hotel successfully")
	return nil
}
