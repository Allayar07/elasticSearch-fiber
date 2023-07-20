package cache

import (
	"context"
	"elasticSearch/internal/models"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	cache *redis.Client
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{
		cache: client,
	}
}

func (c *RedisRepo) Set(key string, value models.Book) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	//if err = c.cache.HMSet(context.Background(), "for searching", bytes).Err(); err != nil {
	//	return err
	//}
	if err = c.cache.Set(context.Background(), key, bytes, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (c *RedisRepo) GetForSearch(value interface{}) (models.Book, error) {
	var book models.Book
	result, err := c.cache.GetSet(context.Background(), "model", value).Bytes()
	if err != nil {
		return models.Book{}, err
	}
	if err = json.Unmarshal(result, &book); err != nil {
		return models.Book{}, err
	}
	fmt.Println(book)
	return book, nil
}
