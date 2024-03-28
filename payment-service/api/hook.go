package apiPayment

import (
	"context"
	"payment-service/internal/model"
	"payment-service/internal/util"
	protoPayment "payment-service/proto/payment"
	protoSdk "payment-service/proto/sdk"

	"github.com/hadanhtuan/go-sdk/common"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhookendpoint"
)

func (bc *PaymentController) InitHook() {
	params := &stripe.WebhookEndpointParams{
		EnabledEvents: []*string{
			stripe.String("charge.succeeded"),
			stripe.String("charge.failed"),
		},
		URL: stripe.String("http://gotrek:3000/api/payment/hook"),
	}
	_, err := webhookendpoint.New(params)

	if err != nil {
		panic("Cannot assign webhook" + err.Error())
	}

}

func (bc *PaymentController) HookPayment(ctx context.Context, req *protoPayment.MsgPaymentIntent) (*protoSdk.BaseResponse, error) {
	filter := &model.PaymentLog{
		BookingId: &req.BookingId,
	}

	result := model.PaymentLogDB.QueryOne(filter, nil)

	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Not found payment log",
		})
	}

	data := result.Data.([]*model.PaymentLog)[0]

	data.Status = req.Status
	data.Event = req.Event
	data.Currency = req.Currency
	data.CanceledAt = req.CanceledAt

	_ = model.PaymentLogDB.Update(filter, data)

	// TODO: send event to property

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Update payment log successfully",
	})
}
