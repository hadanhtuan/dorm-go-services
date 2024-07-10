package main

import (
	"property-service/internal/model"
	"property-service/internal/util"

	"fmt"
	api "property-service/api"
	protoProperty "property-service/proto/property"
	protoUser "property-service/proto/user"

	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/orm"
	"github.com/hadanhtuan/go-sdk/db/redis"
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
	bookingCron := app.NewCronJob()
	bookingCron.SetCronJob(api.InstanceAPI.CheckIfBookingSuccess, 5, 2)

	app.Start()
}

func InitGRPCServer(app *sdk.App) error {
	s := grpc.NewServer()
	UserServiceClient := newUserClient(app)

	api.InitAPI()
	api.InstanceAPI.UserServiceClient = UserServiceClient

	protoProperty.RegisterPropertyServiceServer(s, api.InstanceAPI)

	app.NewGRPCServer(s, app.Config.GRPC.PropertyServiceHost, app.Config.GRPC.PropertyServicePort)

	return nil
}

func newUserClient(app *sdk.App) protoUser.UserServiceClient {
	userServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.UserServiceHost,
		app.Config.GRPC.UserServicePort,
	)
	userConn, err := sdk.NewGRPCClientConn(userServiceHost)
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
