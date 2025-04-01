package models

import "time"

type UsersT struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	BirthDate string    `json:"birth_date"`
	Phone     string    `json:"phone"`
	RoleId    int       `json:"role_id"`
	CiteId    int       `json:"cite_id"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CitiesT struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
