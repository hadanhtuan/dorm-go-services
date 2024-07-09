package main

import (
	"property-service/internal/model"
	"property-service/internal/util"

	"fmt"
	"log"
	"net"
	api "property-service/api"
	protoProperty "property-service/proto/property"
	protoUser "property-service/proto/user"

	"github.com/hadanhtuan/go-sdk/amqp"
	grpcClient "github.com/hadanhtuan/go-sdk/client"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/orm"
	cache "github.com/hadanhtuan/go-sdk/db/redis"
	"gorm.io/gorm"

	"github.com/hadanhtuan/go-sdk"
	"google.golang.org/grpc"
)

func main() {
	config, _ := config.InitConfig(".")
	app := &sdk.App{
		Config: config,
	}

	// Connect Rabbit
	amqp.ConnectRabbit(util.PROPERTY_EXCHANGE, util.PROPERTY_QUEUE, amqp.ExchangeType.Topic)

	// Connect PostgreSQL
	dbOrm := orm.ConnectORM()
	onDBConnected(dbOrm)

	// Connect Redis
	cache.ConnectRedis()

	// Init GRPC
	InitGRPCServer(app)

	//Cronjob
	bookingCron := app.SetupWorker()
	bookingCron.SetCronJob(api.InstanceAPI.CheckIfBookingSuccess, 1, 1)

	app.RunAllCronJob()
}

func InitGRPCServer(app *sdk.App) error {
	propertyServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.PropertyServiceHost,
		app.Config.GRPC.PropertyServicePort,
	)

	lis, err := net.Listen("tcp", propertyServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	UserServiceClient := newUserClient(app)

	api.InitAPI(UserServiceClient)

	protoProperty.RegisterPropertyServiceServer(s, api.InstanceAPI)

	log.Printf("[ Property service ] started on %s", propertyServiceHost)
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}

func newUserClient(app *sdk.App) protoUser.UserServiceClient {
	userServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.UserServiceHost,
		app.Config.GRPC.UserServicePort,
	)
	userConn, err := grpcClient.NewGRPCClientConn(userServiceHost)
	if err != nil {
		return nil
	}

	userServiceClient := protoUser.NewUserServiceClient(userConn)
	return userServiceClient

}

func onDBConnected(db *gorm.DB) {
	model.InitTableAmenity(db)
	model.InitTableProperty(db)
	model.InitTableBooking(db)
	model.InitTableReview(db)
	model.InitTableFavorite(db)
}
