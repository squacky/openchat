package httpapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/squacky/openchat/internal/user/domain"
)

type CreateUser struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
}

func (api *HttpAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUser CreateUser
	err := json.NewDecoder(r.Body).Decode(&createUser)
	if err != nil {
		log.Println("error decoding input body", err)
		api.HandleErr(w, http.StatusBadRequest, err.Error())
		return
	}
	err = api.userService.CreateUser(r.Context(), &domain.User{
		Email:    createUser.Email,
		Phone:    createUser.Phone,
		FullName: createUser.FullName,
	})
	if err != nil {
		api.HandleErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}
