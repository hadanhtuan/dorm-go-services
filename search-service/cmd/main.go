package main

import (
	"search-service/internal"
	"search-service/internal/util"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/elasticsearch"
)

func main() {
	config, _ := config.InitConfig("")

	amqp.ConnectRabbit(util.EXCHANGE, util.QUEUE, amqp.ExchangeType.Topic)
	es.ConnectElasticSearch()

	app := sdk.App{
		Config: config,
	}

	internal.InitGRPCServer(&app)
}
