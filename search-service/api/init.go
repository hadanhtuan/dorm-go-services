package api

import (
	"github.com/hadanhtuan/go-sdk"
	protoSearch "search-service/proto/search"
)

type SearchController struct {
	protoSearch.UnimplementedSearchServiceServer
	App *sdk.App
}
