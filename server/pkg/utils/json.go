package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Response struct {Data any `json:"data"`}
type Error struct {Error string `json:"error"`} // no need for omitempty here; we'll never send an empty error.

// ReadJSON reads a JSON object from an io.ReadCloser,
// closing the reader when it's done.
// It's primarily useful for reading JSON from *http.Request.Body.
func ReadJSON[T any](r io.ReadCloser) (T, error) {
	var v T // declare a variable of type T
	err := json.NewDecoder(r).Decode(&v) // decode the json into v
	return v, errors.Join(err, r.Close()) // close the reader and return any error
}

func DecodeJSON(r io.ReadCloser, v interface{}) error {
	defer r.Close()
	return json.NewDecoder(r).Decode(v)
}

// WriteJSON writes a json object to a http.ResponseWriter, 
// setting the Content-Type header to application/json
func WriteJSON(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application.json")
	return json.NewEncoder(w).Encode(Response{v})
}

// WriteError logs an error, then writes it as a JSON object 
// in the form {"error": <error>}, setting the Content-Type header to application/json
func WriteError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Error{err.Error()})
}