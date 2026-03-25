package models

import "time"

type User struct {
	ID              string    `json:"id" db:"id"`
	FullName        string    `json:"full_name" db:"full_name"`
	Email           string    `json:"email" db:"email"`
	HomeCoordinates string    `json:"home_coordinates,omitempty" db:"home_coordinates"`
	HomeAddress     string    `json:"home_address,omitempty" db:"home_address"`
	Password        string    `json:"-" db:"password"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
