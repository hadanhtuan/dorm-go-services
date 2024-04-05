package apiProperty

import (
	"context"
	"encoding/json"
	"property-service/internal/model"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"

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

func (bc *PropertyController) GetReview(ctx context.Context, req *protoProperty.MsgQueryReview) (*protoSdk.BaseResponse, error) {
	filter := &model.Review{}

	if req.QueryFields.PropertyId != nil && *req.QueryFields.PropertyId != "" {
		filter.PropertyId = req.QueryFields.PropertyId
	}

	result := model.ReviewDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, &orm.QueryOption{
		Order: []string{"created_at desc"},
	})

	data := result.Data.([]*model.Review)

	result = bc.MapReviewWithUser(data)

	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) MapReviewWithUser(data []*model.Review) *common.APIResponse {
	ids := []string{}
	for _, review := range data {
		ids = append(ids, review.UserId)
	}

	result := bc.GetUsers(ids)
	users := make([]*model.User, len(ids))

	encodeData, _ := json.Marshal(result.Data)
	json.Unmarshal(encodeData, &users)

	dict := make(map[string]*model.User, len(ids))
	for _, user := range users {
		dict[user.ID] = user
	}

	for i := 0; i < len(data); i++ {
		data[i].User = *dict[data[i].UserId]
	}

	return &common.APIResponse{
		Data:    data,
		Status:  common.APIStatus.Ok,
		Message: "Query reviews with user successfully",
	}
}
