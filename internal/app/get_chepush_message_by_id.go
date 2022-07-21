package app

import (
	"encoding/json"
	"errors"
	"net/http"
)

type GetChepushMessageByIDRequestBody struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

func (app *Application) GetChepushMessageByID(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	body := &GetChepushMessageByIDRequestBody{}

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return err
	}

	err = app.ChepushilaRepository.SaveMessageFromChepush(r.Context(), body.ID, body.Title, body.Message)
	if err != nil {
		return err
	}

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}
