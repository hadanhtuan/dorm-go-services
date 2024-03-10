package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/model/enum"
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"
	"fmt"

	"github.com/hadanhtuan/go-sdk/common"
)

// Booking
func (bc *BookingController) GetBookingDetail(ctx context.Context, req *protoBooking.MsgGetBookingRequest) (*protoSdk.BaseResponse, error) {

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
	// // 		Status:  common.APIStatus.BadRequest,
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

func (bc *BookingController) GetPropertyDetail(ctx context.Context, req *protoBooking.MsgGetPropertyRequest) (*protoSdk.BaseResponse, error) {
	fmt.Println("aaa")
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}

	result := model.PropertyDB.QueryOne(property)
	if result.Data == nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "GProperty Not Found",
		})
	}
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) GetAllProperty(ctx context.Context, req *protoBooking.MessageQueryRoom) (*protoSdk.BaseResponse, error) {
	result := model.PropertyDB.Query(nil, int(req.Paginate.Offset), int(req.Paginate.Limit))

	result.Message = "Get all properties successfully"
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) CreateProperty(ctx context.Context, req *protoBooking.MsgCreatePropertyRequest) (*protoSdk.BaseResponse, error) {
	propertyType := enum.GetPropertyTypeValue((req.PropertyType))
	property := &model.Property{
		HostId:       req.HostId,
		WardId:       req.WardId,
		NumBeds:      int(req.NumBed),
		NumBedrooms:  int(req.NumBedroom),
		NumBaths:     int(req.NumBath),
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
		PropertyType: &propertyType,
	}

	result := model.PropertyDB.Create(property)
	data := result.Data.([]*model.Property)[0]
	if data != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Ok,
			Message: "Create Property Successfully.",
		})

	}
	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.ServerError,
		Message: "Create Property Failed.",
	})
}
func (bc *BookingController) UpdateProperty(ctx context.Context, req *protoBooking.MsgUpdatePropertyRequest) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}
	propertyType := enum.GetPropertyTypeValue((req.PropertyType))

	propertyUpdated := &model.Property{
		NumBeds:      int(req.NumBed),
		NumBedrooms:  int(req.NumBedroom),
		NumBaths:     int(req.NumBath),
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
		PropertyType: &propertyType,
	}

	result := model.PropertyDB.Update(property, propertyUpdated)
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) DeleteProperty(ctx context.Context, req *protoBooking.MsgDeletePropertyRequest) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}

	result := model.PropertyDB.Delete(property)
	return util.ConvertToGRPC(result)
}

// Review
func (bc *BookingController) CreateReview(ctx context.Context, req *protoBooking.MsgCreateReviewRequest) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		UserId:     req.UserId,
		PropertyId: req.PropertyId,
		ParentId:   req.ParentId,
		Rating:     float64(req.Rating),
		Comment:    req.Comment,
		ImageUrl:   req.ImageUrl,
	}

	result := model.ReviewDB.Create(review)
	data := result.Data.([]*model.Review)[0]
	if data != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Ok,
			Message: "Create Review Successfully.",
		})

	}
	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.ServerError,
		Message: "Create Review Failed.",
	})
}

func (bc *BookingController) UpdateReview(ctx context.Context, req *protoBooking.MsgUpdateReviewRequest) (*protoSdk.BaseResponse, error) {
	reviewId := req.ReviewId
	review := &model.Review{
		ID: reviewId,
	}

	reviewUpdated := &model.Review{
		Rating:   float64(req.Rating),
		Comment:  req.Comment,
		ImageUrl: req.ImageUrl,
	}

	result := model.ReviewDB.Update(review, reviewUpdated)
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) DeleteReview(ctx context.Context, req *protoBooking.MsgDeleteReviewRequest) (*protoSdk.BaseResponse, error) {
	reviewId := req.ReviewId
	review := &model.Review{
		ID: reviewId,
	}

	result := model.ReviewDB.Delete(review)
	return util.ConvertToGRPC(result)
}

func (bc *BookingController) GetReview(ctx context.Context, req *protoBooking.MessageQueryReview) (*protoSdk.BaseResponse, error) {
	result := model.ReviewDB.Query(req.QueryFields.Id, 1, 1)
	result.Message = "Get all reviews successfully"
	return util.ConvertToGRPC(result)

}
