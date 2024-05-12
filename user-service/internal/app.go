package internal

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	api "user-service/api"
	userProto "user-service/proto/user"
	pkg "github.com/hadanhtuan/go-sdk"
)

func InitGRPCServer(app *pkg.App) error {
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
	userProto.RegisterUserServiceServer(s, &api.UserController{})

	log.Printf("User Service started on %s", userServiceHost)

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
