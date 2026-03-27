package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tanyaP0405/go-grpc-graphql-micro/account"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Printf("failed to connect to database: %v", err)
		}
		return
	})
	defer r.Close()
	log.Println("connected to database")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
