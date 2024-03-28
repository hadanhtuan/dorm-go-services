package api

import (
	protoSearch "search-service/proto/search"
	"github.com/hadanhtuan/go-sdk"
)

type SearchController struct {
	protoSearch.UnimplementedSearchServiceServer
	App *sdk.App
}

func (pc *SearchController) InitController() {
	pc.InitIndex()
	pc.InitRoutingAMQP()
}
