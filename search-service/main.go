package main

import (
	"fmt"
	"log"
	"net"
	"search-service/internal/util"
	"search-service/api"
	searchProto "search-service/proto/search"
	pkg "github.com/hadanhtuan/go-sdk"
	"google.golang.org/grpc"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/elasticsearch"
	cache "github.com/hadanhtuan/go-sdk/db/redis"
)

func main() {
	config, _ := config.InitConfig(".")

	app := sdk.App{
		Config: config,
	}

	// Connect Rabbit
	amqp.ConnectRabbit(util.SEARCH_EXCHANGE, util.SEARCH_QUEUE, amqp.ExchangeType.Topic)

	// Connect Redis
	cache.ConnectRedis()

	// Connect Elasticsearch
	es.ConnectElasticSearch()

	// Init GRPC
	InitGRPCServer(&app)
}

func InitGRPCServer(app *pkg.App) error {
	searchServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.SearchServiceHost,
		app.Config.GRPC.SearchServicePort,
	)
	lis, err := net.Listen("tcp", searchServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	api.InitAPI(app)

	searchProto.RegisterSearchServiceServer(s, api.InstanceAPI)

	log.Printf("[ Search service ] started on %s", searchServiceHost)

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
