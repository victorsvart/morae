package roomhandler

import (
	"encoding/json"
	"errors"
	"morae/internal/dto/roomdto"
	"morae/internal/usecase/room"
	"morae/internal/utils"
	"net/http"
)

type RoomHandler struct {
	Usecases *room.RoomUsecases
}


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

func (rh *RoomHandler) GetRoomUserId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, errors.New("Id is invalid"))
		return
	}

	response, err := rh.Usecases.GetById.Execute(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

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

func (rh *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  if id == "" {
    utils.RespondWithError(w, http.StatusBadRequest, errors.New("Id is empty"))
    return
  }

  if err := rh.Usecases.DeleteRoom.Execute(r.Context(), id); err != nil {
    utils.RespondWithError(w, http.StatusBadRequest, err)
    return
  }

  utils.RespondWithSuccess(w, http.StatusOK, "Deleted succesfully")
}
