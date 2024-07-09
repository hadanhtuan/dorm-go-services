package api

import (
	"encoding/json"
	"fmt"
	"search-service/internal/model"
	"search-service/internal/util"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	es "github.com/hadanhtuan/go-sdk/db/elasticsearch"
	cache "github.com/hadanhtuan/go-sdk/db/redis"
)

func (sc *SearchAPI) SaveSearchRecord(searchText string, userId *string) {
	documentId := strings.Trim(searchText, " ")
	documentId = strings.ReplaceAll(documentId, " ", "_")

	if userId != nil && *userId != "" {
		key := fmt.Sprintf("%s:%s", util.CacheSearchTracking, *userId) // key = search:tracking:user1
		go sc.CacheRecently(searchText, key)                          // save to redis
	}

	document := &model.SearchTrackingDocument{
		SearchText:  searchText,
		SearchCount: 1,
	}

	sc.UpsertSearchDocument(documentId, document) // update tracking index
}

func (sc *SearchAPI) CacheRecently(searchText, key string) {
	recently := []string{}

	cache.Get(key, &recently)

	recently = append([]string{searchText}, recently...)
	recently = util.UniqueSliceElements(recently)

	if len(recently) > util.MaxSuggestion {
		recently = recently[:util.MaxSuggestion]
	}

	cache.Set(key, recently, 10*24*time.Hour)
}

func (sc *SearchAPI) GetRecent(userId *string, size int) []string {
	recently := []string{}
	if userId != nil && *userId != "" {
		prefix := fmt.Sprintf("%s:%s", util.CacheSearchTracking, *userId)
		cache.Get(prefix, &recently)

		if len(recently) > size {
			recently = recently[:size]
		}
	}

	return recently
}


func (sc *SearchAPI) UpsertSearchDocument(documentId string, document *model.SearchTrackingDocument) {
	script := `ctx._source.searchCount = ctx._source.searchCount != null ? ctx._source.searchCount += 1 : 1`
	dataDoc, _ := json.Marshal(document)

	es.UpdateDocument(util.TrackingIndex, documentId, &update.Request{
		Script: types.InlineScript{
			Source: script,
		},
		Upsert: json.RawMessage(dataDoc),
	})
}
