package apiBooking

import (
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hadanhtuan/go-sdk/common"
)

func (pc *BookingController) GetBookingDetail(ctx context.Context, req *protoBooking.MsgGetBookingRequest) (*protoSdk.BaseResponse, error) {

	// filter := map[string]interface{}{}

	fmt.Print("GetBookingDetail", req)
	// filter["email"] = req.Email

	// result := model.UserDB.QueryOne(filter)
	// if result.Data == nil {
	// 	return &protoSdk.BaseResponse{
	// 		Status:  common.APIStatus.Unauthorized,
	// 		Message: "Username or password incorrect",
	// 	}, nil
	// }
	// data := result.Data.([]*model.User)[0]

	// isVerify := sdk.VerifyPassword(req.Password, data.Password)
	// if !isVerify || result.Data == nil {
	// 	return &protoSdk.BaseResponse{
	// 		Status:  common.APIStatus.Unauthorized,
	// 		Message: "Username or password incorrect",
	// 	}, nil
	// }

	// jwtPayload := &common.JWTPayload{
	// 	ID:       data.ID,
	// 	Email:    data.Email,
	// 	DeviceID: req.DeviceId,
	// }
	// token, err := aws.NewJWT(jwtPayload)

	// if err != nil {
	// 	return &protoSdk.BaseResponse{
	// 		Status:  common.APIStatus.BadRequest,
	// 		Message: "Error generate JWT. Error Detail: " + err.Error(),
	// 	}, nil
	// }

	// //TODO: save login log
	// loginLog := &model.LoginLog{
	// 	UserId:    data.ID,
	// 	UserAgent: req.UserAgent,
	// 	IpAddress: req.IpAddress,
	// 	DeviceID:  req.DeviceId,
	// }
	// model.LoginLogDB.Create(loginLog)

	encodeData, _ := json.Marshal("OK")
	return &protoSdk.BaseResponse{
		Status:  common.APIStatus.Ok,
		Message: "AAAAA",
		Data:    string(encodeData),
	}, nil
}
