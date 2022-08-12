package api

import (
	"CaseMajoo/repository"
	"encoding/json"
	"net/http"
)

type UserListSuccessResponse struct {
	Users []repository.User `json:"users"`
}

type UserErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) userList(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	response := UserListSuccessResponse{}
	response.Users = make([]repository.User, 0)

	users, err := api.usersRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(UserErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, user := range users {
		response.Users = append(response.Users, repository.User{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Password: user.Password,
		})
	}

	encoder.Encode(response)
}
