package internal

import (
	"fmt"
	grpcClient "github.com/hadanhtuan/go-sdk/client"
	"log"
	"net"
	api "property-service/api"
	protoProperty "property-service/proto/property"
	protoUser "property-service/proto/user"

	"github.com/hadanhtuan/go-sdk"
	"google.golang.org/grpc"
)

func InitGRPCServer(app *sdk.App) error {
	propertyServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.PaymentServiceHost,
		app.Config.GRPC.PropertyServicePort,
	)

	lis, err := net.Listen("tcp", propertyServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	newApi := &api.PropertyController{}
	protoProperty.RegisterPropertyServiceServer(s, newApi)
	
	newApi.UserServiceClient = newUserClient(app)
	newApi.InitController()

	log.Printf("Property server started on %s", propertyServiceHost)
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
