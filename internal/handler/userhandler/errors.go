// Package userhandler contains handlers for managing user-related HTTP requests.
package userhandler

import "errors"

// ErrInvalidID indicates the user ID is invalid.
var ErrInvalidID = errors.New("user ID is invalid")
