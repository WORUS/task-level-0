package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"task-level-0/internal/api/handler"
	"task-level-0/internal/api/stan/publisher"
	"task-level-0/internal/api/stan/subscriber"
	"task-level-0/internal/repository"
	"task-level-0/internal/server"
	"task-level-0/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

const (
	cacheParseBase       = 10
	cacheParseBitSize    = 32
	cacheDefaultCapacity = 100
)

//TODO delete cmd/ sub

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variabless: %s", err.Error())
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	cache := make(map[string][]byte)
	cacheCapacity := cacheDefaultCapacity
	temp, err := strconv.ParseInt(os.Getenv("CACHE_CAPACITY"), cacheParseBase, cacheParseBitSize)
	if err != nil {
		logrus.WithError(err).Infof("error with parse cache size cause set default_size = %d", cacheDefaultCapacity)
	} else {
		cacheCapacity = int(temp)
	}

	repository := repository.NewRepository(dbpool, cacheCapacity, cache)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	srv := new(server.Server)

	repository.RestoreCache()

	go func() {
		if err := srv.Run(os.Getenv("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	sc, err := stan.Connect(os.Getenv("CLUSTER_ID"), os.Getenv("CLIENT_ID"))
	if err != nil {
		fmt.Print("error")
	}

	sub := subscriber.NewSubscriber(sc, os.Getenv("NATS_SUBJECT"), service)
	sub.Run()

	pub := publisher.NewPublisher(sc, os.Getenv("NATS_SUBJECT"), true)
	go pub.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	pub.Stop()
	sub.Stop()

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}
