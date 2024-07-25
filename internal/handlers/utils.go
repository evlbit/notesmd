package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func parseJSON(r *http.Request, payload any) error {
	err := json.NewDecoder(r.Body).Decode(payload)

	if err != nil && err.Error() == "EOF" {
		return fmt.Errorf("request body missing")
	} else if err != nil {
		return fmt.Errorf("invalid request body json")
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, res any) {
	w.Header().Add("Content-Type", "application/json")

	if res == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		jsonRes, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "something went wrong"})
			return
		}

		w.WriteHeader(status)
		w.Write(jsonRes)
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

func writeInternalError(w http.ResponseWriter) {
	writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
}
