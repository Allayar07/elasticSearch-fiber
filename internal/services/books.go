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
	cache    repository.Cache
	kafka    repository.Queue
	esClient *elasticsearch.Client
}

func NewBookService(repo repository.Books, cache repository.Cache, kafka repository.Queue, client *elasticsearch.Client) *BookService {
	return &BookService{
		repos:    repo,
		cache:    cache,
		kafka:    kafka,
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
	if err = s.cache.Set("model", book); err != nil {
		return 0, err
	}
	if err = s.kafka.PublishTopic("books", book.Name); err != nil {
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

func (s *BookService) GetFormCache(search interface{}) (models.Book, error) {
	return s.cache.GetForSearch(search)
}

func (s *BookService) Sync() error {
	books, err := s.repos.GetForSync()
	if err != nil {
		return err
	}
	return SyncWithDB(books, s.esClient)
}
