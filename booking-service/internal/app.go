package internal

import (
	api "booking-service/api"
	bookingProto "booking-service/proto/booking"
	"fmt"
	"log"
	"net"

	pkg "github.com/hadanhtuan/go-sdk"
	"google.golang.org/grpc"
)


func InitGRPCServer(app *pkg.App) error {
	bookingServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.BookingServiceHost,
		app.Config.GRPC.BookingServicePort,
	)
	lis, err := net.Listen("tcp", bookingServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	bookingProto.RegisterBookingServiceServer(s, &api.BookingController{})
	log.Printf("gRPC Server started on %s", bookingServiceHost)

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server down")

	return nil
}
