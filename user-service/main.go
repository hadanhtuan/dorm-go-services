package main

import (
	"fmt"
	"user-service/internal"
	"user-service/internal/model"

	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/aws"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig(".")
	aws.ConnectAWS()
	dbOrm := orm.ConnectDB()
	app := sdk.App{
		Config: config,
	}

	onDBConnected(dbOrm)
	internal.InitGRPCServer(&app)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableUser(db)
	model.InitTableLoginSession(db)
}
