package main

import (
	"fmt"
	. "github.com/hadanhtuan/go-sdk"
	config "github.com/hadanhtuan/go-sdk/config"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
	"user-service/internal"
	database "user-service/internal/db"
	"user-service/internal/model"
)

func main() {
	config, _ := config.InitConfig("")
	dbOrm := orm.Connect(config.DBOrm)
	app := App{
		Config: config,
		DBOrm:  dbOrm,
	}

	onDBConnected(dbOrm)
	internal.InitGRPCServer(&app)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	database.Migrate(db)
	model.InitTableUser(db)
}
