package books

import (
	"fmt"
	"go-library-service/entities"
	"log"
)

type Service interface {
	GetBooks() ([]entities.Book, error)
	GetBook(string) ([]entities.Book, error)
	PostBook(book entities.Book) ([]entities.Book, error)
	DeleteBook(string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (sr *service) GetBooks() ([]entities.Book, error) {
	books := []entities.Book{}
	book := entities.Book{"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"}
	books = append(books, book)

	return books, nil
}

func (sr *service) GetBook(uuid string) ([]entities.Book, error) {
	books := []entities.Book{}
	book := entities.Book{uuid}
	books = append(books, book)

	return books, nil
}

func (sr *service) PostBook(book entities.Book) ([]entities.Book, error) {
	log.Print(fmt.Sprintf("The book %s has been posted!", book.Uuid))
	books := []entities.Book{}
	books = append(books, book)

	return books, nil
}

func (sr *service) DeleteBook(uuid string) error {
	log.Print(fmt.Sprintf("The book %s has been deleted!", uuid))
	return nil
}
