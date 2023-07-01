package app

import (
	"elasticSearch/internal/configs"
	"elasticSearch/internal/handler"
	"elasticSearch/internal/repository"
	"elasticSearch/internal/services"
	database "elasticSearch/pkg/db"
	"elasticSearch/pkg/elasticSearch_client"
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
	repos := repository.NewRepository(db)
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
