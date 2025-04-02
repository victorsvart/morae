package userdomain

import (
	"time"
)

type User struct {
	ID           uint64
	FullName     string
	EmailAddress EmailAddress
	Password     Password
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) UserChecksAndSets(email, password string) error {
	emailAddress, err := NewEmailAddress(email)
	if err != nil {
		return err
	}

	pass := &Password{Value: password}
	pass.HashPassword()

	u.EmailAddress = *emailAddress
	u.Password = *pass
	return nil
}
