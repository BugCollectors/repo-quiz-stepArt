package app

import (
	"encoding/json"
	"errors"
	"net/http"
)

type PostBody struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

func (app *Application) Post(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	body := &PostBody{}

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return err
	}

	err = app.ChepushilaRepository.WriteMessageFromChepush(r.Context(), body.ID, body.Title, body.Message)
	if err != nil {
		return err
	}

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}
