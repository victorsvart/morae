package usermapper

import (
	"morae/internal/domain/userdomain"
	"morae/internal/store/postgres"
)

func ToDomain(user *postgres.UserEntity) *userdomain.User {
	return &userdomain.User{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: *userdomain.SetEmail(user.EmailAddress),
		Password:     *userdomain.SetPassword(user.Password),
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func ToEntity(user *userdomain.User) *postgres.UserEntity {
	return &postgres.UserEntity{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: user.EmailAddress.Value,
		Password:     user.Password.Value,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func FromDto(input *userdomain.UserDto) (*userdomain.User, error) {
	return &userdomain.User{
		ID:           input.ID,
		FullName:     input.FullName,
		EmailAddress: *userdomain.SetEmail(input.EmailAddress),
		Password:     *userdomain.SetPassword(input.Password),
	}, nil
}

func FromInput(input *userdomain.UserInput) (*userdomain.User, error) {
	user := &userdomain.User{
		FullName:     input.FullName,
		EmailAddress: userdomain.EmailAddress{},
		Password:     userdomain.Password{},
	}
	user.UserChecksAndSets(input.EmailAddress, input.Password)

	return user, nil
}

func ToResponse(user *postgres.UserEntity) *userdomain.UserResponse {
	return &userdomain.UserResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: user.EmailAddress,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
