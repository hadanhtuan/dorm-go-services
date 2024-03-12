package api

import (
	"context"
	"search-service/internal/util"
	"github.com/hadanhtuan/go-sdk/common"
	protoSdk "search-service/proto/sdk"
	protoSearch "search-service/proto/search"
)

func (pc *SearchController) Search(ctx context.Context, req *protoSearch.MsgSearch) (*protoSdk.BaseResponse, error) {

	return util.ConvertToGRPC(&common.APIResponse{
		Status: common.APIStatus.Ok,
	})
}
