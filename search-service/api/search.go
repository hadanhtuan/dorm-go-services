package api

import (
	"context"
	es "search-service/internal/elasticsearch"
	"search-service/internal/util"
	protoSdk "search-service/proto/sdk"
	protoSearch "search-service/proto/search"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hadanhtuan/go-sdk/common"
)

func (pc *SearchController) Search(ctx context.Context, req *protoSearch.MsgSearchProperty) (*protoSdk.BaseResponse, error) {
	size := int(req.Paginate.Limit)
	from := (int(req.Paginate.Offset) - 1) * size

	query := &search.Request{
		From:  &from,
		Size:  &size,
		Sort:  []types.SortCombinations{},
		Query: &types.Query{},
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status: common.APIStatus.Ok,
	})
}

func (pc *SearchController) InitIndex() {
	for key, value := range util.IndicesMap {
		es.CreateIndex(key, value)
	}
}
