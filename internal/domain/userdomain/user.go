package userdomain

import (
	"time"
)

// User represents a domain model for a system user.
type User struct {
	ID           uint64
	FullName     string
	EmailAddress EmailAddress
	Password     Password
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// SetCredentials validates and sets the user's email and password.
func (u *User) SetCredentials(email, password string) error {
	emailAddress, err := NewEmailAddress(email)
	if err != nil {
		return err
	}

	pass := &Password{Value: password}
	if err := pass.HashPassword(); err != nil {
		return err
	}

	u.EmailAddress = *emailAddress
	u.Password = *pass
	return nil
}
