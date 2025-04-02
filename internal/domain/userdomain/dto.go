package userdomain

import (
	"errors"
	"time"
)

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

func (ui *UserInput) Validate() error {
	if ui.FullName == "" {
		return errors.New("full name is required")
	}

	if ui.EmailAddress == "" {
		return errors.New("email address is required")
	}
	if ui.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type UserResponse struct {
	ID           uint64    `json:"id"`
	FullName     string    `json:"fullName"`
	EmailAddress string    `json:"emailAddress"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updateAt"`
}
