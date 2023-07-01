package services

import (
	"context"
	"elasticSearch/internal/models"
	"elasticSearch/internal/repository"
	"github.com/elastic/go-elasticsearch/v8"
	"strconv"
)

type BookService struct {
	repos    repository.Books
	esClient *elasticsearch.Client
}

func NewBookService(repo repository.Books, client *elasticsearch.Client) *BookService {
	return &BookService{
		repos:    repo,
		esClient: client,
	}
}

func (s *BookService) CreateBook(book models.Book) (int, error) {
	id, err := s.repos.Create(book)
	if err != nil {
		return 0, err
	}
	bookId := strconv.Itoa(id)
	book.Id = id
	if err = IndexingToElasticSearch(context.Background(), bookId, book, s.esClient); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *BookService) Search(searchInput string) ([]models.Book, error) {
	books, err := SearchBook(searchInput, s.esClient)
	if err != nil {
		return []models.Book{}, err
	}
	return books, nil
}

func (s *BookService) Delete(ids models.DeleteIds) error {
	if err := s.repos.DeleteBooks(ids); err != nil {
		return err
	}
	if err := DeleteFromElastic(ids, s.esClient); err != nil {
		return err
	}
	return nil
}

func (s *BookService) Update(book models.Book) error {
	if err := s.repos.UpdateBook(book); err != nil {
		return err
	}
	if err := UpdateFRomElastic(book, s.esClient); err != nil {
		return err
	}
	return nil
}
