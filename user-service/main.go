package main

import (
	"fmt"
	api "user-service/api"
	"user-service/internal/model"
	userProto "user-service/proto/user"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/orm"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	config, _ := config.InitConfig(".")

	// aws.ConnectAWS()

	app := &sdk.App{
		Config: config,
	}

	// Connect PostgreSQL
	dbOrm := orm.ConnectORM()
	onDBConnected(dbOrm)

	// Init GRPC
	InitGRPCServer(app)

	app.Start()
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableUser(db)
	model.InitTableLoginSession(db)
}

func InitGRPCServer(app *sdk.App) error {
	s := grpc.NewServer()

	api.InitAPI()
	userProto.RegisterUserServiceServer(s, api.InstanceAPI)

	app.NewGRPCServer(s, app.Config.GRPC.UserServiceHost, app.Config.GRPC.UserServicePort)

	return nil
}
