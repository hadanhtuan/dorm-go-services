package apiBooking

import (
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"

	"github.com/hadanhtuan/go-sdk/common"
)

// Booking
func (bc *BookingController) GetBookingDetail(ctx context.Context, req *protoBooking.MsgGetBooking) (*protoSdk.BaseResponse, error) {

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Get Booking Detail Successfully.",
	})
}
