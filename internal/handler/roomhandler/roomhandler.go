// Package roomhandler contains HTTP handlers for room-related operations.
package roomhandler

import (
	"encoding/json"
	"errors"
	"morae/internal/dto/roomdto"
	"morae/internal/usecase/room"
	"morae/internal/utils"
	"net/http"
)

// RoomHandler holds the use cases for room operations.
type RoomHandler struct {
	Usecases *room.Usecases
}

// GetAllRooms handles the request to retrieve all rooms with optional pagination and filtering.
func (rh *RoomHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	var input roomdto.GetRoomPaged
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	response, err := rh.Usecases.GetAllRooms.Execute(r.Context(), &input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetRoomUserID handles the request to get a room by its ID.
func (rh *RoomHandler) GetRoomUserID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, errors.New("id is invalid"))
		return
	}

	response, err := rh.Usecases.GetByID.Execute(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// CreateRoom handles the request to create a new room.
func (rh *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var input roomdto.RoomInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	response, err := rh.Usecases.CreateRoom.Execute(r.Context(), &input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// UpdateRoom handles the request to update an existing room.
func (rh *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var input roomdto.RoomDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}

	response, err := rh.Usecases.UpdateRoom.Execute(r.Context(), &input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// DeleteRoom handles the request to delete a room by ID.
func (rh *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, errors.New("id is empty"))
		return
	}

	if err := rh.Usecases.DeleteRoom.Execute(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, "deleted successfully")
}
