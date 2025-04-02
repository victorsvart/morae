package userdomain

import "time"

type UserDto struct {
	ID           uint64 `json:"id"`
	FullName     string `json:"fullName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type UserInput struct {
	FullName     string `json:"fullName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type UserResponse struct {
	ID           uint64    `json:"id"`
	FullName     string    `json:"fullName"`
	EmailAddress string    `json:"emailAddress"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updateAt"`
}
