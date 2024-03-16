package main

import (
	"fmt"
	"search-service/internal"
	es "search-service/internal/elasticsearch"
	"search-service/internal/util"

	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/aws"
	config "github.com/hadanhtuan/go-sdk/config"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig("")

	aws.ConnectAWS()
	amqp.ConnectRabbit(util.EXCHANGE, util.QUEUE, amqp.ExchangeType.Topic)
	es.ConnectElasticSearch()
	dbOrm := orm.ConnectDB()
	onDBConnected(dbOrm)

	app := sdk.App{
		Config: config,
	}

	internal.InitGRPCServer(&app)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	// model.InitTableUser(db)
	// model.InitTableLoginLog(db)
}
