package internal

import (
	"fmt"
	"log"
	"net"
	"search-service/api"
	searchProto "search-service/proto/search"

	pkg "github.com/hadanhtuan/go-sdk"
	"google.golang.org/grpc"
)

func InitGRPCServer(app *pkg.App) error {
	searchServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.SearchServiceHost,
		app.Config.GRPC.SearchServicePort,
	)
	lis, err := net.Listen("tcp", searchServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	newApi := &api.SearchController{App: app}
	searchProto.RegisterSearchServiceServer(s, newApi)

	log.Printf("Search server started on %s", searchServiceHost)

	newApi.InitController()

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
