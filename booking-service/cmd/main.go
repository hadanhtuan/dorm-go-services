package main

import (
	"booking-service/internal"
	"booking-service/internal/model"
	"booking-service/internal/util"

	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig("")

	amqp.ConnectRabbit(util.EXCHANGE, util.QUEUE, amqp.ExchangeType.Topic)
	dbOrm := orm.ConnectDB()
	onDBConnected(dbOrm)

	app := &sdk.App{
		Config: config,
	}

	internal.InitGRPCServer(app)
}

func onDBConnected(db *gorm.DB) {
	model.InitTableAmenity(db)
	model.InitTableProperty(db)
	model.InitTableBooking(db)
	model.InitTableReview(db)
}
