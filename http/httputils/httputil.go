package httputils

import (
	"encoding/json"
	"net/http"
	"reviews/errdefs"
)

type ApiHandlerFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string) error

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	body, err := json.Marshal(v)
	if err != nil {
		return errdefs.NewStatusError(500, err)
	}

	_, err = w.Write(body)
	return errdefs.NewStatusError(500, err)
}

func ReadJSON(r *http.Request, model interface{}) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		return nil, errdefs.NewStatusError(500, err)
	}

	defer r.Body.Close()

	return model, nil
}
