package apiProperty

import (
	"context"
	"property-service/internal/model"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
)

func (bc *PropertyAPI) CreateFavorite(ctx context.Context, req *protoProperty.MsgFavorite) (*protoSdk.BaseResponse, error) {
	favorite := &model.Favorite{
		UserId:     req.UserId,
		PropertyId: req.PropertyId,
	}

	result := model.FavoriteDB.Create(favorite)

	return util.ConvertToGRPC(result)
}

func (bc *PropertyAPI) DeleteFavorite(ctx context.Context, req *protoProperty.MsgId) (*protoSdk.BaseResponse, error) {
	favorite := &model.Favorite{
		ID: req.Id,
	}

	result := model.FavoriteDB.Delete(favorite)
	return util.ConvertToGRPC(result)
}

func (bc *PropertyAPI) GetFavorite(ctx context.Context, req *protoProperty.MsgQueryFavorite) (*protoSdk.BaseResponse, error) {
	filter := &model.Favorite{}

	if req.QueryFields.Id != "" {
		filter.ID = req.QueryFields.Id
	}

	if req.QueryFields.UserId != "" {
		filter.UserId = req.QueryFields.UserId
	}

	result := model.FavoriteDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, nil)

	return util.ConvertToGRPC(result)
}
