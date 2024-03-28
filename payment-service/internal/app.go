package internal

import (
	"fmt"
	"log"
	"net"
	api "payment-service/api"
	paymentProto "payment-service/proto/payment"

	pkg "github.com/hadanhtuan/go-sdk"
	"google.golang.org/grpc"
)

func InitGRPCServer(app *pkg.App) error {
	paymentServiceHost := fmt.Sprintf(
		"%s:%s",
		app.Config.GRPC.PaymentServiceHost,
		app.Config.GRPC.PaymentServicePort,
	)
	lis, err := net.Listen("tcp", paymentServiceHost)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	newApi := &api.PaymentController{}
	paymentProto.RegisterPaymentServiceServer(s, newApi)
	log.Printf("Payment server started on %s", paymentServiceHost)

	newApi.InitController()

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
