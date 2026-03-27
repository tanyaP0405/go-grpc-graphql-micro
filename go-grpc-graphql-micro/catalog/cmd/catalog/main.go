package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tanyaP0405/go-grpc-graphql-micro/catalog"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Connected to database")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, 8080))
}
