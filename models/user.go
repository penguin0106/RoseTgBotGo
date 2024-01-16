package models

import "time"

// User структура для хранения информации о пользователе в базе данных MongoDB.
type User struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PhoneNumber  string    `json:"phone_number"`
	Role         string    `json:"role"` // "admin" или "user"
	Registration time.Time `json:"registration_date"`
}
