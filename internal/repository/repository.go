package repository

import (
	"elasticSearch/internal/models"
	"elasticSearch/internal/repository/cache"
	"elasticSearch/internal/repository/kafka_repo"
	"elasticSearch/internal/repository/postgres"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Books interface {
	Create(book models.Book) (int, error)
	DeleteBooks(ids models.DeleteIds) error
	UpdateBook(book models.Book) error
}

type Cache interface {
	Set(key string, value models.Book) error
	GetForSearch(value interface{}) (models.Book, error)
}

type Queue interface {
	PublishTopic(topic, value string) error
}
type Repository struct {
	Books
	Cache
	Queue
}

func NewRepository(db *pgxpool.Pool, client *redis.Client, kafka *kafka.Producer) *Repository {
	return &Repository{
		Books: postgres.NewBooksRepo(db),
		Cache: cache.NewRedisRepo(client),
		Queue: kafka_repo.NewKafkaRepo(kafka),
	}
}
