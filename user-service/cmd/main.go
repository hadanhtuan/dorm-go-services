package main

import (
	"fmt"
	"user-service/internal"
	"user-service/internal/model"

	. "github.com/hadanhtuan/go-sdk"
	orm "github.com/hadanhtuan/go-sdk/db/orm"

	config "github.com/hadanhtuan/go-sdk/config"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig("")
	dbOrm := orm.Connect(config.DBOrm)
	app := App{
		Config: config,
		// DBOrm:  dbOrm,
	}

	internal.InitGRPC(&app)
	internal.InitRoute(&app)
	onDBConnected(dbOrm)
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableUser(db)
}
