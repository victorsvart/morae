// Package userdto defines the Data Transfer Objects for user operations.
package userdto

import (
	"time"
)

// UserDto represents a full user data structure used internally or in responses.
type UserDto struct {
	ID           uint64 `json:"id"`
	FullName     string `json:"fullName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

// UserInput represents the expected input payload for creating or updating a user.
type UserInput struct {
	FullName     string `json:"fullName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

// Validate checks whether all required fields in UserInput are provided.
func (ui *UserInput) Validate() error {
	if ui.FullName == "" {
		return ErrFullNameIsRequired
	}
	if ui.EmailAddress == "" {
		return ErrEmailIsRequired
	}
	if ui.Password == "" {
		return ErrPasswordIsRequired
	}
	return nil
}

// UserResponse represents the user structure returned in API responses.
type UserResponse struct {
	ID           uint64    `json:"id"`
	FullName     string    `json:"fullName"`
	EmailAddress string    `json:"emailAddress"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updateAt"`
}
