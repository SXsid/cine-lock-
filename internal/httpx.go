package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Reponse struct {
	Status bool   `json:"status"`
	Data   any    `json:"data,omitempty"`
	Msg    string `json:"message,omitempty"`
}

func WriteJson(w http.ResponseWriter, data any, status int, msg string) {
	value := Reponse{
		Status: true,
		Data:   data,
	}
	json, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "go")
	w.WriteHeader(status)
	if _, err := w.Write(json); err != nil {
		fmt.Println(err)
	}
}

func ReadJson[T any](r *http.Request) (T, error) {
	var res T
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}
