package books

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-library-service/entities"
	"net/http"
)

type Response struct {
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Books []entities.Book
}

type Request struct {
	Data RequestData `json:"data"`
}

type RequestData struct {
	Book entities.Book `json:"Book"`
}

type Endpoints struct {
	GetBooks   endpoint.Endpoint
	GetBook    endpoint.Endpoint
	PostBook   endpoint.Endpoint
	DeleteBook endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetBooks:   MakeGetBooks(s),
		GetBook:    MakeGetBook(s),
		PostBook:   MakePostBook(s),
		DeleteBook: MakeDeleteBook(s),
	}
}

func MakeGetBooks(sr Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		books, err := sr.GetBooks()

		return Response{ResponseData{books}}, err
	}
}

func MakeGetBook(sr Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(map[string]string)
		if !ok {
			panic(http.StatusText(http.StatusBadRequest))
		}
		uuid := r["uuid"]
		books, err := sr.GetBook(uuid)

		return Response{ResponseData{books}}, err
	}
}

func MakePostBook(sr Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(Request)
		book := entities.Book{r.Data.Book.Uuid}
		if !ok {
			panic(http.StatusText(http.StatusBadRequest))
		}
		books, err := sr.PostBook(book)

		return Response{ResponseData{books}}, err
	}
}

func MakeDeleteBook(sr Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(map[string]string)
		if !ok {
			panic(http.StatusText(http.StatusBadRequest))
		}
		err := sr.DeleteBook(r["uuid"])

		return nil, err
	}
}
