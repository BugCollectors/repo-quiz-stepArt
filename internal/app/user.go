package app

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"study/internal/types"
)

func (h *Handler) createUser(_ http.ResponseWriter, r *http.Request) {
	var user types.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error decode post body: %v", err))
	}

	id, errUser := h.services.CreateUser(user)
	if errUser != nil {
		newErrorResponse(fmt.Sprintf("error create user: %v", errUser))
	}

	fmt.Printf("Create User id=%d\n", id)
}

func (h *Handler) getUser(_ http.ResponseWriter, r *http.Request) {
	//id, ok := r.Context().Value("id").(string)
	//if !ok {
	//}

	var id struct {
		ID string `json:"id" db:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error decode post body: %v", err))
	}

	user, errUser := h.services.GetUser(id)
	if errUser != nil {
		newErrorResponse(fmt.Sprintf("error get user: %v", err))
	}

	logrus.Infof(fmt.Sprintf("User: %v", user))

}
