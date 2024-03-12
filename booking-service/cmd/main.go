package main

import (
	"booking-service/internal"
	"booking-service/internal/model"
	"booking-service/internal/util"
	"fmt"

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
	dbOrm := orm.ConnectDB()

	app := &sdk.App{
		Config: config,
	}

	onDBConnected(dbOrm)
	internal.InitGRPCServer(app)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableAmenity(db)
	model.InitTableBooking(db)
	model.InitTablePropertyAmenity(db)
	model.InitTableProperty(db)
	model.InitTableReview(db)
}
