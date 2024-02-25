package main

import (
	"fmt"
	"user-service/internal"
	"user-service/internal/model"
	aws "user-service/internal/aws"

	sdk "github.com/hadanhtuan/go-sdk"
	config "github.com/hadanhtuan/go-sdk/config"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig("")
	dbOrm := orm.Connect(config.DBOrm)
	app := sdk.App{
		Config: config,
		DBOrm:  dbOrm,
	}

	aws.TestKMS()
	onDBConnected(dbOrm)
	internal.InitGRPCServer(&app)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableUser(db)
	model.InitTableLoginLog(db)
}
