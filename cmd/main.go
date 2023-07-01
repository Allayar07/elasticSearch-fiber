package main

import (
	"elasticSearch/internal/app"
	"elasticSearch/internal/configs"
	"log"
)

func main() {
	config, err := configs.LoadConfiguration("./configs/config.yml")
	if err != nil {
		log.Println(err)
	}

	if err = app.InitApp(config); err != nil {
		log.Fatalln(err)
	}
}
