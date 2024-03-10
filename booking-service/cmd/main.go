package main

import (
	"booking-service/internal"
	"booking-service/internal/model"
	"fmt"

	"github.com/hadanhtuan/go-sdk"
	aws "github.com/hadanhtuan/go-sdk/aws"
	config "github.com/hadanhtuan/go-sdk/config"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig("")
	dbOrm := orm.Connect(config.DBOrm)
	aws.ConnectAWS()
	app := sdk.App{
		Config: config,
		DBOrm:  dbOrm,
	}

	onDBConnected(dbOrm)
	internal.InitGRPCServer(&app)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableBooking(db)
	model.InitTableProperty(db)
	model.InitTableReview(db)
}
