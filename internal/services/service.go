package services

import (
	"elasticSearch/internal/models"
	"elasticSearch/internal/repository"
	"github.com/elastic/go-elasticsearch/v8"
)

type Books interface {
	CreateBook(book models.Book) (int, error)
	Search(searchInput string) ([]models.Book, error)
	Delete(ids models.DeleteIds) error
	Update(book models.Book) error
	GetFormCache(search interface{}) (models.Book, error)
	Sync() error
}

type Service struct {
	Books
}

func NewService(repo *repository.Repository, EsClient *elasticsearch.Client) *Service {
	return &Service{
		Books: NewBookService(repo.Books, repo.Cache, repo.Queue, EsClient),
	}
}
