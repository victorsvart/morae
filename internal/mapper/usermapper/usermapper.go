// Package usermapper provides mapping functions between user domain models,
// DTOs, and persistence entities.
package usermapper

import (
	"morae/internal/domain/userdomain"
	"morae/internal/dto/userdto"
	"morae/internal/store/postgres"
)

// ToDomain converts a UserEntity to a User domain model.
func ToDomain(user *postgres.UserEntity) *userdomain.User {
	return &userdomain.User{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: userdomain.EmailAddress{Value: user.EmailAddress},
		Password:     userdomain.Password{Value: user.Password},
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

// ToEntity converts a User domain model to a UserEntity.
func ToEntity(user *userdomain.User) postgres.UserEntity {
	return postgres.UserEntity{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: user.EmailAddress.Value,
		Password:     user.Password.Value,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

// FromDto maps a UserDto to a User domain model.
func FromDto(input *userdto.UserDto) (*userdomain.User, error) {
	return &userdomain.User{
		ID:           input.ID,
		FullName:     input.FullName,
		EmailAddress: userdomain.EmailAddress{Value: input.EmailAddress},
		Password:     userdomain.Password{Value: input.Password},
	}, nil
}

// FromInput maps a UserInput to a User domain model and applies validation/transformation.
func FromInput(input *userdto.UserInput) (userdomain.User, error) {
	user := userdomain.User{
		FullName:     input.FullName,
		EmailAddress: userdomain.EmailAddress{},
		Password:     userdomain.Password{},
	}

	err := user.SetCredentials(input.EmailAddress, input.Password)
	if err != nil {
		return userdomain.User{}, err
	}

	return user, nil
}

// ToResponse maps a UserEntity to a UserResponse DTO.
func ToResponse(user *postgres.UserEntity) userdto.UserResponse {
	return userdto.UserResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: user.EmailAddress,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
