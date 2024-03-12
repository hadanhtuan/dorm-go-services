package main

import (
	"fmt"
	"user-service/internal"
	"user-service/internal/model"

	"github.com/hadanhtuan/go-sdk"
	aws "github.com/hadanhtuan/go-sdk/aws"
	config "github.com/hadanhtuan/go-sdk/config"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig("")
	dbOrm := orm.ConnectDB()
	aws.ConnectAWS()
	app := sdk.App{
		Config: config,
	}

	onDBConnected(dbOrm)
	internal.InitGRPCServer(&app)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableUser(db)
	model.InitTableLoginLog(db)
}
