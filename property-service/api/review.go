package apiProperty

import (
	"property-service/internal/model"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
	"context"
	"fmt"

	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (bc *PropertyController) CreateReview(ctx context.Context, req *protoProperty.MsgCreateReview) (*protoSdk.BaseResponse, error) {
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

func (bc *PropertyController) UpdateReview(ctx context.Context, req *protoProperty.MsgUpdateReview) (*protoSdk.BaseResponse, error) {
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

func (bc *PropertyController) DeleteReview(ctx context.Context, req *protoProperty.MsgDeleteReview) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		ID: req.ReviewId,
	}

	result := model.ReviewDB.Delete(review)
	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) GetReview(ctx context.Context, req *protoProperty.MessageQueryReview) (*protoSdk.BaseResponse, error) {
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
