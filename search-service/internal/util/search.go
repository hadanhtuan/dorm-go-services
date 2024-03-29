package util

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// constant do not have memory address(cannot use &)
var (
	customAnalyzer   = "custom_analyzer"
	customNormalizer = "custom_normalizer"

	CacheSearchDocument = "search:document"
	CacheSearchTracking = "search:tracking" //search and save trend word

	PropertyIndex = "property_index"
	CityIndex     = "city_index"
)

var IndicesMap = map[string]*create.Request{
	PropertyIndex: PropertyIndexCnf,
	CityIndex:     CityIndexCnf,
}

var PropertyIndexCnf = &create.Request{
	Settings: &types.IndexSettings{
		Analysis: &types.IndexSettingsAnalysis{
			Analyzer: map[string]types.Analyzer{
				customAnalyzer: &types.CustomAnalyzer{
					Tokenizer: "standard",
					Filter:    []string{"lowercase", "classic"},
				},
			},
			Normalizer: map[string]types.Normalizer{
				customNormalizer: &types.LowercaseNormalizer{
					Type: "lowercase",
				},
			},
		},
	},
	Mappings: &types.TypeMapping{
		Properties: map[string]types.Property{
			"id": types.KeywordProperty{
				Type: "keyword",
			},
			"createdAt": types.DateProperty{
				Type: "date",
			},
			"amenities": types.NestedProperty{
				Type: "nested",
				Properties: map[string]types.Property{
					"id": types.KeywordProperty{
						Type: "keyword",
					},
					"name": types.KeywordProperty{
						Type:       "keyword",
						Normalizer: &customNormalizer,
					},
				},
			},
			"bookings": types.NestedProperty{
				Type: "nested",
				Properties: map[string]types.Property{
					"id": types.KeywordProperty{
						Type: "keyword",
					},
					"checkInDate": types.IntegerNumberProperty{
						Type: "integer",
					},
					"checkoutDate": types.IntegerNumberProperty{
						Type: "integer",
					},
					"status": types.KeywordProperty{
						Type:       "keyword",
						Normalizer: &customNormalizer,
					},
				},
			},
			"propertyType": types.KeywordProperty{
				Type: "keyword",
			},
			"overallRate": types.FloatNumberProperty{
				Type: "float",
			},
			"maxGuests": types.IntegerNumberProperty{
				Type: "integer",
			},
			"maxPets": types.IntegerNumberProperty{
				Type: "integer",
			},
			"numGuests": types.IntegerNumberProperty{
				Type: "integer",
			},
			"numBeds": types.IntegerNumberProperty{
				Type: "integer",
			},
			"numBedrooms": types.IntegerNumberProperty{
				Type: "integer",
			},
			"numBathrooms": types.IntegerNumberProperty{
				Type: "integer",
			},
			"isGuestFavor": types.BooleanProperty{
				Type: "boolean",
			},
			"isAllowPet": types.BooleanProperty{
				Type: "boolean",
			},
			"isSelfCheckIn": types.BooleanProperty{
				Type: "boolean",
			},
			"isInstantBook": types.BooleanProperty{
				Type: "boolean",
			},
			"title": types.TextProperty{
				Type:     "text",
				Analyzer: &customAnalyzer,
			},
			"body": types.TextProperty{
				Type:     "text",
				Analyzer: &customAnalyzer,
			},
			"cityCode": types.KeywordProperty{
				Type: "keyword",
			},
			"nationCode": types.KeywordProperty{
				Type: "keyword",
			},
			"lat": types.KeywordProperty{
				Type: "keyword",
			},
			"long": types.KeywordProperty{
				Type: "keyword",
			},
			"nightPrice": types.FloatNumberProperty{
				Type: "float",
			},
			"serviceFee": types.FloatNumberProperty{
				Type: "float",
			},
		},
	},
}

var CityIndexCnf = &create.Request{
	Settings: &types.IndexSettings{
		Analysis: &types.IndexSettingsAnalysis{
			Normalizer: map[string]types.Normalizer{
				customNormalizer: &types.LowercaseNormalizer{
					Type: "lowercase",
				},
			},
		},
	},
	Mappings: &types.TypeMapping{
		Properties: map[string]types.Property{
			"cityCode": types.KeywordProperty{
				Type:       "keyword",
				Normalizer: &customNormalizer,
			},
			"searchCount": types.IntegerNumberProperty{
				Type: "integer",
			},
		},
	},
}
