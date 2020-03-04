package infrastructure

import (
	"context"
	"encoding/json"
	"go-library-service/books"
	"net/http"

	thttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(r *mux.Router, e books.Endpoints) http.Handler {
	options := []thttp.ServerOption{
		thttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/books").Handler(thttp.NewServer(
		e.GetBooks,
		decodeGetBooks,
		getEncodeResponse,
		options...,
	))

	r.Methods("GET").Path("/books/{uuid}").Handler(thttp.NewServer(
		e.GetBook,
		decodeGetBook,
		getEncodeResponse,
		options...,
	))

	r.Methods("POST").Path("/books").Handler(thttp.NewServer(
		e.PostBook,
		decodePostBookRequest,
		postEncodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/books/{uuid}").Handler(thttp.NewServer(
		e.DeleteBook,
		decodeDeleteXlsxRequest,
		deleteEncodeResponse,
		options...,
	))

	return r
}

func decodeGetBooks(_ context.Context, r *http.Request) (interface{}, error) {
	q := mux.Vars(r)

	return q, nil
}

func decodeGetBook(_ context.Context, r *http.Request) (interface{}, error) {
	q := mux.Vars(r)

	return q, nil
}

func decodePostBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request books.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeDeleteXlsxRequest(_ context.Context, r *http.Request) (interface{}, error) {
	q := mux.Vars(r)

	return q, nil
}

func getEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(response)
}

func postEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(response)
}

func deleteEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)

	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
