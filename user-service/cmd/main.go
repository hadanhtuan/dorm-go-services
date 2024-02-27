package main

import (
	"fmt"
	"user-service/internal"
	"user-service/internal/model"
	"gorm.io/gorm"
	"github.com/hadanhtuan/go-sdk"
	aws "github.com/hadanhtuan/go-sdk/aws"
	config "github.com/hadanhtuan/go-sdk/config"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
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
	model.InitTableUser(db)
	model.InitTableLoginLog(db)
}
