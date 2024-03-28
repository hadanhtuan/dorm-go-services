package internal

import (
	api "property-service/api"
	propertyProto "property-service/proto/property"
	"fmt"
	"log"
	"net"

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
	propertyProto.RegisterPropertyServiceServer(s, &api.PropertyController{})
	log.Printf("Property server started on %s", propertyServiceHost)

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
