package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/model/enum"
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"

	"github.com/hadanhtuan/go-sdk/common"
)

// Booking
func (bc *BookingController) GetBookingDetail(ctx context.Context, req *protoBooking.MsgGetBooking) (*protoSdk.BaseResponse, error) {

	// filter := map[string]interface{}{}

	// fmt.Print("GetBookingDetail", req)
	// filter["id"] = req.BookingId
	// id, ok := filter["id"]
	// if !ok {
	// 	// Xử lý lỗi nếu không thể ép kiểu giá trị của trường "id" thành uuid.UUID
	// }
	// var booking = model.Booking{
	// 	ID: string(id),
	// }

	// result := model.BookingDB.QueryOne(&booking)
	// fmt.Print("GetBookingDetail", result.Data)

	// if result.Data == nil {
	// 	return &protoSdk.BaseResponse{
	// 		Status:  common.APIStatus.NotFound,
	// 		Message: "Booking Not Found",
	// 	}, nil
	// }
	// data := result.Data.([]*model.Booking)[0]

	// // isVerify := sdk.VerifyPassword(req.Password, data.Password)
	// // if !isVerify || result.Data == nil {
	// // 	return &protoSdk.BaseResponse{
	// // 		Status:  common.APIStatus.Unauthorized,
	// // 		Message: "Username or password incorrect",
	// // 	}, nil
	// // }

	// // jwtPayload := &common.JWTPayload{
	// // 	ID:       data.ID,
	// // 	Email:    data.Email,
	// // 	DeviceID: req.DeviceId,
	// // }
	// // token, err := aws.NewJWT(jwtPayload)

	// // if err != nil {
	// // 	return &protoSdk.BaseResponse{
	// // 		Status:  common.APIStatus.Bad,
	// // 		Message: "Error generate JWT. Error Detail: " + err.Error(),
	// // 	}, nil
	// // }

	// // //TODO: save login log
	// // loginLog := &model.LoginLog{
	// // 	UserId:    data.ID,
	// // 	UserAgent: req.UserAgent,
	// // 	IpAddress: req.IpAddress,
	// // 	DeviceID:  req.DeviceId,
	// // }
	// // model.LoginLogDB.Create(loginLog)

	// encodeData, _ := json.Marshal(data)
	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Get Booking Detail Successfully.",
	})
}

//Property

func (bc *BookingController) GetPropertyDetail(ctx context.Context, req *protoBooking.MsgGetProperty) (*protoSdk.BaseResponse, error) {
	property := &model.Property{
		ID: req.PropertyId,
	}

	result := model.PropertyDB.QueryOne(property)
	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Property Not Found",
		})
	}
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) GetAllProperty(ctx context.Context, req *protoBooking.MsgQueryProperty) (*protoSdk.BaseResponse, error) {
	result := model.PropertyDB.Query(nil, req.Paginate.Offset, req.Paginate.Limit)

	result.Message = "Get all properties successfully"
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) CreateProperty(ctx context.Context, req *protoBooking.MsgCreateProperty) (*protoSdk.BaseResponse, error) {
	propertyType := enum.PropertyTypeValue(req.PropertyType)
	property := &model.Property{
		HostId:       req.HostId,
		WardId:       req.WardId,
		NumBeds:      req.NumBed,
		NumBedrooms:  req.NumBedroom,
		NumBaths:     req.NumBath,
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
		PropertyType: &propertyType,
	}

	result := model.PropertyDB.Create(property)
	if result.Status != common.APIStatus.Created {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Create Property Failed.",
		})
	}
	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Create Property Successfully.",
	})
}

func (bc *BookingController) UpdateProperty(ctx context.Context, req *protoBooking.MsgUpdateProperty) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}
	propertyType := enum.PropertyTypeValue(req.PropertyType)

	propertyUpdated := &model.Property{
		NumBeds:      req.NumBed,
		NumBedrooms:  req.NumBedroom,
		NumBaths:     req.NumBath,
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
		PropertyType: &propertyType,
	}

	result := model.PropertyDB.Update(property, propertyUpdated)
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) DeleteProperty(ctx context.Context, req *protoBooking.MsgDeleteProperty) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}

	result := model.PropertyDB.Delete(property)
	return util.ConvertToGRPC(result)
}