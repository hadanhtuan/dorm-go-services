package api

import (
	"context"
	"fmt"
	"net"
	"search-service/internal/model"
	"search-service/internal/util"
	protoSdk "search-service/proto/sdk"
	protoSearch "search-service/proto/search"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

	fmt.Println("what the fuck", size)

	queryField := req.QueryFields

	mustQuery := []types.Query{}
	shouldQuery := []types.Query{}

	if queryField.Amenities != nil && len(queryField.Amenities) > 0 {
		amenityQuery := []types.Query{}
		ignoreUnmapped := true

		for _, amenity := range queryField.Amenities {
			fmt.Println(*amenity.Id)
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

	if queryField.Title != nil {
		// go sc.SaveSearchRecord(*queryField.Title, *queryField.UserId)

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

	fmt.Println("result", result.Total)

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
	size := int(req.Paginate.Limit)
	from := (int(req.Paginate.Offset) - 1) * size

	query := &search.Request{
		Size: &size,
		From: &from,
		Sort: []types.SortCombinations{},
	}

	result := es.Search[model.Property](util.PropertyIndex, query)
	return util.ConvertToGRPC(result)
}

func (sc *SearchController) ListPropertyByIP(ctx context.Context, req *protoSearch.MsgIP) (*protoSdk.BaseResponse, error) {
	size := int(req.Paginate.Limit)
	from := (int(req.Paginate.Offset) - 1) * size

	mustQuery := []types.Query{}

	// countryCode, err := sc.GetCountryByIp(req.IpAddress)
	// fmt.Println(countryCode)

	// if err != nil {
	// 	return util.ConvertToGRPC(&common.APIResponse{
	// 		Status:  common.APIStatus.BadRequest,
	// 		Message: "Error parsing country. Error detail: " + err.Error(),
	// 	})
	// }

	mustQuery = append(mustQuery, types.Query{
		Match: map[string]types.MatchQuery{
			"nationCode": {Query: "USA"},
		},
	})

	query := &search.Request{
		From: &from,
		Size: &size,
		Sort: []types.SortCombinations{},
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must: mustQuery,
			},
		},
	}

	result := es.Search[model.Property](util.PropertyIndex, query)
	return util.ConvertToGRPC(result)
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
