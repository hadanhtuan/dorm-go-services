package api

import (
	protoSearch "search-service/proto/search"
	"github.com/hadanhtuan/go-sdk"
)

type SearchAPI struct {
	protoSearch.UnimplementedSearchServiceServer
	App *sdk.App
}

func InitAPI(App *sdk.App) {
	sa := &SearchAPI{
		App: App,
	}

	sa.InitIndex()
	sa.InitRoutingAMQP()
	InstanceAPI = sa
}

var (
	InstanceAPI *SearchAPI
)
