package main

import (
	"fmt"
	"log"
	"net"
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
}

func onDBConnected(db *gorm.DB) {
	fmt.Println("Connected to DB " + db.Name())
	model.InitTableUser(db)
	model.InitTableLoginSession(db)
}

func InitGRPCServer(app *sdk.App) error {
	userServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.UserServiceHost,
		app.Config.GRPC.UserServicePort,
	)
	lis, err := net.Listen("tcp", userServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	api.InitAPI()
	userProto.RegisterUserServiceServer(s, api.InstanceAPI)

	log.Printf("[ User service ] started on %s", userServiceHost)

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
