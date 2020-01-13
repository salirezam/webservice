package controllers

import "net/http"

import "encoding/json"

func RegisterController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application-json")
	enc := json.NewEncoder(w)
	_ = enc.Encode(data)
}
