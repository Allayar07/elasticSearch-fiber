package app

import (
	"elasticSearch/internal/configs"
	"elasticSearch/internal/handler"
	"elasticSearch/internal/repository"
	"elasticSearch/internal/services"
	database "elasticSearch/pkg/db"
	"elasticSearch/pkg/elasticSearch_client"
	"elasticSearch/pkg/kafka_producer"
	"elasticSearch/pkg/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func InitApp(cfg *configs.Configs) error {
	db, err := database.NewPostgres(cfg)
	if err != nil {
		return err
	}

	elasticClient, err := elasticSearch_client.SetUpElasticSearch()
	if err != nil {
		return err
	}

	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		return err
	}
	producerKafka, err := kafka_producer.ProducerQueues()
	if err != nil {
		return err
	}

	defer func() {
		producerKafka.Close()
		db.Close()
		_ = redisClient.Close()
	}()
	repos := repository.NewRepository(db, redisClient, producerKafka)
	service := services.NewService(repos, elasticClient)
	apiHandler := handler.NewHandler(service)
	app := apiHandler.InitRoutes()

	go func() {
		if err = app.Listen(":7575"); err != nil {
			log.Fatalln(err)
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	<-signals

	if err = app.Shutdown(); err != nil {
		log.Fatalln(err)
	}

	return nil
}
