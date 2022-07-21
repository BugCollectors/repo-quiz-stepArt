package app_layer

import (
	"encoding/json"
	"errors"
	"net/http"
)

type GetChepushMessageByIDRequestBody struct {
	ID string `json:"id,omitempty"`
}

func (app *Application) GetChepushMessageByID(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	body := &GetChepushMessageByIDRequestBody{}

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return err
	}

	err = app.ChepushilaRepository.GetChepushMessageByID(r.Context(), body.ID)
	if err != nil {
		return err
	}

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}
