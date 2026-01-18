package dto

type PeopleCreate struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	PlaceOfBirth string `json:"place_of_birth" binding:"required"`
	DateOfBirth  string `json:"date_of_birth" binding:"required"`
	Address      string `json:"address" binding:"required"`
}

type PeoplePatch struct {
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	PlaceOfBirth *string `json:"place_of_birth,omitempty"`
	DateOfBirth  *string `json:"date_of_birth,omitempty"`
	Address      *string `json:"address,omitempty"`
}
