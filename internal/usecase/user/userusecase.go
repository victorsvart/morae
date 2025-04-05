package user

import (
	"morae/internal/store/postgres"
)

// Usecases groups all user-related use cases.
type Usecases struct {
	GetByID GetUserByIDUsecase // Retrieves a user by ID.
	Create  CreateUserUsecase  // Creates a new user.
	Update  UpdateUserUsecase  // Updates an existing user.
	Delete  DeleteUserUsecase  // Deletes a user by ID.
}

// NewUserUsecases initializes and returns a Usecases struct with the given repository.
func NewUserUsecases(repo postgres.UserRepository) *Usecases {
	return &Usecases{
		GetByID: &GetUserByID{repo},
		Create:  &Create{repo},
		Update:  &Update{repo},
		Delete:  &Delete{repo},
	}
}
