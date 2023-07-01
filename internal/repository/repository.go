package repository

import (
	"elasticSearch/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Books interface {
	Create(book models.Book) (int, error)
	DeleteBooks(ids models.DeleteIds) error
	UpdateBook(book models.Book) error
}

type Repository struct {
	Books
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Books: NewBooksRepo(db),
	}
}
