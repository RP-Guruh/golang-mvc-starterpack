package models

import "time"

type People struct {
	ID           uint      `json:"id"`
	FirstName    string    `json:"first_name" binding:"required"`
	LastName     string    `json:"last_name" binding:"required"`
	PlaceOfBirth string    `json:"place_of_birth" binding:"required"`
	DateOfBirth  time.Time `json:"date_of_birth" binding:"required"`
	Address      string    `json:"address" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
