package internal

import (
	"encoding/json"
	"net/http"
)

type Reponse struct {
	Data any    `json:"data,omitempty"`
	Msg  string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error    string            `json:"error"`
	FieldErr map[string]string `json:"fields,omitempty"`
}

func write(w http.ResponseWriter, v any, status int) {
	res, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "go")
	w.WriteHeader(status)
	w.Write(res)
}

func writeOK(w http.ResponseWriter, msg string, status int) {
	v := Reponse{
		Msg: msg,
	}
	write(w, v, status)
}

func WriteJSON(w http.ResponseWriter, data any, status int) {
	value := Reponse{
		Data: data,
	}
	write(w, value, status)
}

func WriteError(w http.ResponseWriter, msg string, status int) {
	v := ErrorResponse{
		Error: msg,
	}
	write(w, v, status)
}

func WriteValidation(w http.ResponseWriter, Fielderr map[string]string) {
	v := ErrorResponse{
		Error:    "error while validating",
		FieldErr: Fielderr,
	}
	write(w, v, http.StatusBadRequest)
}

func ReadJson[T any](r *http.Request) (T, error) {
	var res T
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}
