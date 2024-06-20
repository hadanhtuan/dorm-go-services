package api

import (
	"context"
	"fmt"
	"net"
	"search-service/internal/model"
	"search-service/internal/util"
	protoSdk "search-service/proto/sdk"
	protoSearch "search-service/proto/search"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/hadanhtuan/go-sdk/common"
	es "github.com/hadanhtuan/go-sdk/db/elasticsearch"
	"github.com/ipinfo/go/v2/ipinfo"
)

func (sc *SearchController) InitIndex() {
	for key, value := range util.IndicesMap {
		es.CreateIndex(key, value)
	}
}

func (sc *SearchController) SearchProperty(ctx context.Context, req *protoSearch.MsgSearchProperty) (*protoSdk.BaseResponse, error) {
	size := int(req.Paginate.Limit)
	from := (int(req.Paginate.Offset) - 1) * size

	queryField := req.QueryFields

	mustQuery := []types.Query{}
	shouldQuery := []types.Query{}

	if queryField.Amenities != nil && len(queryField.Amenities) > 0 {
		amenityQuery := []types.Query{}
		ignoreUnmapped := true

		for _, amenity := range queryField.Amenities {
			amenityQuery = append(amenityQuery, types.Query{
				Nested: &types.NestedQuery{
					Path:           "amenities",
					IgnoreUnmapped: &ignoreUnmapped,
					Query: &types.Query{
						Bool: &types.BoolQuery{
							Must: []types.Query{
								{
									Term: map[string]types.TermQuery{
										"amenities.id": {
											Value: *amenity.Id,
										},
									},
								},
							},
						},
					},
				},
			})
		}
		mustQuery = append(mustQuery, amenityQuery...)
	}

	if queryField.NightPriceMin != nil && queryField.NightPriceMax != nil {
		priceMin := types.Float64(*queryField.NightPriceMin)
		priceMax := types.Float64(*queryField.NightPriceMax)
		mustQuery = append(mustQuery, types.Query{
			Range: map[string]types.RangeQuery{
				"nightPrice": types.NumberRangeQuery{
					Gte: &priceMin,
					Lte: &priceMax,
				},
			},
		})
	}

	if queryField.MaxGuests != nil {
		maxGuests := types.Float64(*queryField.MaxGuests)
		mustQuery = append(mustQuery, types.Query{
			Range: map[string]types.RangeQuery{
				"maxGuests": types.NumberRangeQuery{
					Gte: &maxGuests,
				},
			},
		})
	}

	if queryField.MaxPets != nil {
		maxPets := types.Float64(*queryField.MaxPets)
		mustQuery = append(mustQuery, types.Query{
			Range: map[string]types.RangeQuery{
				"maxPets": types.NumberRangeQuery{
					Gte: &maxPets,
				},
			},
		})
	}

	if queryField.NumBedrooms != nil {
		if queryField.NumBedrooms != nil {
			mustQuery = append(mustQuery, types.Query{
				Term: map[string]types.TermQuery{
					"numBedrooms": {Value: *queryField.NumBedrooms},
				},
			})
		}
	}

	if queryField.NumBeds != nil {
		if queryField.NumBeds != nil {
			mustQuery = append(mustQuery, types.Query{
				Term: map[string]types.TermQuery{
					"numBeds": {Value: *queryField.NumBeds},
				},
			})
		}
	}

	if queryField.NumBathrooms != nil {
		if queryField.NumBathrooms != nil {
			mustQuery = append(mustQuery, types.Query{
				Term: map[string]types.TermQuery{
					"numBathrooms": {Value: *queryField.NumBathrooms},
				},
			})
		}
	}

	if queryField.IsAllowPet != nil {
		mustQuery = append(mustQuery, types.Query{
			Term: map[string]types.TermQuery{
				"isAllowPet": {Value: *queryField.IsAllowPet},
			},
		})
	}

	if queryField.IsGuestFavor != nil {
		mustQuery = append(mustQuery, types.Query{
			Term: map[string]types.TermQuery{
				"isGuestFavor": {Value: *queryField.IsGuestFavor},
			},
		})
	}

	if queryField.IsInstantBook != nil {
		mustQuery = append(mustQuery, types.Query{
			Term: map[string]types.TermQuery{
				"isInstantBook": {Value: *queryField.IsInstantBook},
			},
		})
	}

	if queryField.IsSelfCheckIn != nil {
		mustQuery = append(mustQuery, types.Query{
			Term: map[string]types.TermQuery{
				"isSelfCheckIn": {Value: *queryField.IsSelfCheckIn},
			},
		})
	}

	if queryField.CityCode != nil {
		mustQuery = append(mustQuery, types.Query{
			Term: map[string]types.TermQuery{
				"cityCode": {Value: *queryField.CityCode},
			},
		})
	}

	// Keyword have standard analyzer(lowercase filter). Use match instead of term
	if queryField.NationCode != nil {
		mustQuery = append(mustQuery, types.Query{
			Match: map[string]types.MatchQuery{
				"nationCode": {Query: *queryField.NationCode},
			},
		})
	}

	if queryField.CityCode != nil {
		mustQuery = append(mustQuery, types.Query{
			Match: map[string]types.MatchQuery{
				"cityCode": {Query: *queryField.CityCode},
			},
		})
	}

	if queryField.PropertyType != nil {
		mustQuery = append(mustQuery, types.Query{
			Match: map[string]types.MatchQuery{
				"propertyType": {Query: *queryField.PropertyType},
			},
		})
	}

	if queryField.Status != nil {
		mustQuery = append(mustQuery, types.Query{
			Match: map[string]types.MatchQuery{
				"status": {Query: *queryField.Status},
			},
		})
	}

	if queryField.HostName != nil && *queryField.HostName != "" {
		mustQuery = append(mustQuery, types.Query{
			MatchPhrasePrefix: map[string]types.MatchPhrasePrefixQuery{
				"hostName": {Query: *queryField.HostName},
			},
		})
	}

	if queryField.Title != nil && *queryField.Title != "" {
		go sc.SaveSearchRecord(*queryField.Title, queryField.UserId)

		shouldQuery = append(shouldQuery, types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Query: *queryField.Title,
				Fields: []string{
					"title^2",
					"body",
				},
				Fuzziness: "AUTO",
			},
		})
	}

	query := &search.Request{
		From: &from,
		Size: &size,
		Sort: []types.SortCombinations{},
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must:   mustQuery,
				Should: shouldQuery,
			},
		},
	}

	result := es.Search[model.Property](util.PropertyIndex, query)

	return util.ConvertToGRPC(result)
}

func (sc *SearchController) GetNation(ctx context.Context, req *protoSearch.MsgIP) (*protoSdk.BaseResponse, error) {
	size := 0
	field := "nationCode"

	aggQuery := map[string]types.Aggregations{
		"count_distinct": {
			Terms: &types.TermsAggregation{
				Field: &field,
			},
		},
	}

	query := &search.Request{
		Size:         &size,
		Sort:         []types.SortCombinations{},
		Aggregations: aggQuery,
	}

	conn := es.GetConnection()

	res, err := conn.Client.Search().Index(util.PropertyIndex).Request(query).Do(context.Background())

	if err != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Total:   0,
			Message: "Error query index " + util.PropertyIndex + ". Error detail: " + err.Error(),
			Status:  common.APIStatus.BadRequest,
		})
	}

	data := res.Aggregations["count_distinct"].(*types.StringTermsAggregate)
	bucket := data.Buckets.([]types.StringTermsBucket)

	nations := []string{}

	for _, v := range bucket {
		nations = append(nations, v.Key.(string))
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Data:   nations,
		Status: common.APIStatus.Ok,
	})
}

func (sc *SearchController) RenderSuggestion(ctx context.Context, req *protoSearch.MsgSuggestion) (*protoSdk.BaseResponse, error) {
	var wg sync.WaitGroup

	type SuggestionResponse struct {
		Popular  []string `json:"popular"`
		Recently []string `json:"recently"`
	}

	popular := []string(nil)
	recently := []string(nil)

	size := int(req.Paginate.Limit)
	from := (int(req.Paginate.Offset) - 1) * size

	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		popular, _ = sc.GetPopular(&from, &size)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		recently = sc.GetRecent(req.UserId, size)
	}(&wg)

	wg.Wait()

	res := SuggestionResponse{
		Popular:  popular,
		Recently: recently,
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Data:   res,
		Status: common.APIStatus.Ok,
	})

}

func (sc *SearchController) SearchTitlePrefix(ctx context.Context, req *protoSearch.MessageSearchPrefix) (*protoSdk.BaseResponse, error) {

	size := int(req.Paginate.Limit)
	from := (int(req.Paginate.Offset) - 1) * size

	query := &search.Request{
		From: &from,
		Size: &size,
		Sort: []types.SortCombinations{
			types.SortOptions{
				SortOptions: map[string]types.FieldSort{
					"searchCount": {
						Order: &sortorder.SortOrder{
							Name: util.Desc,
						},
					},
				},
			},
		},
		Query: &types.Query{
			Prefix: map[string]types.PrefixQuery{
				"searchText": {
					Value: req.SearchText,
				},
			},
		},
		Highlight: &types.Highlight{
			Fields: map[string]types.HighlightField{
				"searchText": {},
			},
		},
	}
	result := es.Search[model.SearchTrackingDocument](
		util.TrackingIndex,
		query,
	)

	if result.Status != common.APIStatus.Ok {
		return util.ConvertToGRPC(result)
	}

	data := result.Data.([]model.SearchTrackingDocument)

	for i := 0; i < len(data); i++ {
		key := strings.ReplaceAll(strings.Trim(data[i].SearchText, " "), " ", "_")
		data[i].Highlight = result.Headers[key]
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   data,
	})
}

func (sc *SearchController) GetPopular(from, size *int) ([]string, error) {
	query := &search.Request{
		From: from,
		Size: size,
		Sort: []types.SortCombinations{
			types.SortOptions{
				SortOptions: map[string]types.FieldSort{
					"searchCount": {
						Order: &sortorder.SortOrder{
							Name: util.Desc,
						},
					},
				},
			},
		},
	}

	result := es.Search[model.SearchTrackingDocument](
		util.TrackingIndex,
		query,
	)

	if result.Status != common.APIStatus.Ok {
		return nil, fmt.Errorf(result.Message)
	}

	data := result.Data.([]model.SearchTrackingDocument)

	popularly := []string(nil)
	for _, record := range data {
		popularly = append(popularly, record.SearchText)
	}

	return popularly, nil
}

func (sc *SearchController) GetCountryByIp(ip string) (string, error) {
	const token = "33e32a0a50947a"
	client := ipinfo.NewClient(nil, nil, token)
	info, err := client.GetIPInfo(net.ParseIP(ip))

	if err != nil {
		return "", err
	}

	return info.Country, nil
}
