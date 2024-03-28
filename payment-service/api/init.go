package apiPayment

import (
	protoPayment "payment-service/proto/payment"
)

type PaymentController struct {
	protoPayment.UnimplementedPaymentServiceServer
}

func (bc *PaymentController) InitController() {
	// bc.InitHook()
}