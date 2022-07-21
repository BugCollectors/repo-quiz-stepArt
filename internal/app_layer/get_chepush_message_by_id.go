package app_layer

import (
	"encoding/json"
	"errors"
	"net/http"
)

type GetChepushMessageByIDRequestBody struct {
	ID string `json:"id,omitempty"`
}

type GetChepushMessageByIDResponse struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

func (app *Application) GetChepushMessageByID(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	body := &GetChepushMessageByIDRequestBody{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return err
	}

	message, err := app.ChepushilaRepository.GetChepushMessageByID(r.Context(), body.ID)
	if err != nil {
		return err
	}

	response := GetChepushMessageByIDResponse{
		ID:      message.ID,
		Title:   message.Title,
		Message: message.Message,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}
