package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"
	"fmt"

	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (bc *BookingController) CreateReview(ctx context.Context, req *protoBooking.MsgCreateReview) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		UserId:     req.UserId,
		PropertyId: req.PropertyId,
		Rating:     req.Rating,
		Comment:    req.Comment,
		ImageUrl:   req.ImageUrl,
	}

	if req.ParentId != "" {
		review.ParentId = &req.ParentId
	}

	result := model.ReviewDB.Create(review)
	if result.Status != common.APIStatus.Created {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Create Review Failed.",
		})
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Create Review Successfully.",
	})
}

func (bc *BookingController) UpdateReview(ctx context.Context, req *protoBooking.MsgUpdateReview) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		ID: req.ReviewId,
	}

	reviewUpdated := &model.Review{
		Rating:   req.Rating,
		Comment:  req.Comment,
		ImageUrl: req.ImageUrl,
	}

	result := model.ReviewDB.Update(review, reviewUpdated)
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) DeleteReview(ctx context.Context, req *protoBooking.MsgDeleteReview) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		ID: req.ReviewId,
	}

	result := model.ReviewDB.Delete(review)
	return util.ConvertToGRPC(result)
}

func (bc *BookingController) GetReview(ctx context.Context, req *protoBooking.MessageQueryReview) (*protoSdk.BaseResponse, error) {
	filter := &model.Review{}

	if req.QueryFields.PropertyId != nil {
		filter.PropertyId = req.QueryFields.PropertyId
	}

	result := model.ReviewDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, &orm.QueryOption{
		Order: []string{"created_at asc"},
	})

	data := result.Data.([]*model.Review)
	fmt.Println(data)

	return util.ConvertToGRPC(result)
}
