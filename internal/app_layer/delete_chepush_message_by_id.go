package app_layer

import (
	"errors"
	"net/http"
)

func (s *Application) DeleteChepushMessageByID(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}