package main

import (
	"payment-service/internal"
	"payment-service/internal/model"
	"payment-service/internal/util"

	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/orm"
	"github.com/stripe/stripe-go/v76"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig(".")

	amqp.ConnectRabbit(util.PAYMENT_EXCHANGE, util.PAYMENT_QUEUE, amqp.ExchangeType.Topic)
	stripe.Key = config.Stripe.SecretKey

	dbOrm := orm.ConnectDB()
	onDBConnected(dbOrm)

	app := &sdk.App{
		Config: config,
	}

	internal.InitGRPCServer(app)
}

func onDBConnected(db *gorm.DB) {
	model.InitTablePaymentLog(db)
}
