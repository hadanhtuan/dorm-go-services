package es

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// constant do not have memory address(cannot use &)
var (
	emptyReplacement = " "
	customAnalyzer   = "custom_analyzer"
	normalizer       = "custom_normalizer"

	CacheSearchDocument = "search:document"
	CacheSearchTracking = "search:tracking"

	SearchAll  = "ALL"
	SearchPost = "POST"
	SearchUser = "USER"

	PropertyIndices       = "property_index"
	UserIndices           = "user_index"
	SearchTrackingIndices = "search_tracking_index"
)

var PropertyIndicesCnf = &create.Request{
	Settings: &types.IndexSettings{
		Analysis: &types.IndexSettingsAnalysis{
			Analyzer: map[string]types.Analyzer{
				"custom_analyzer": &types.CustomAnalyzer{
					Tokenizer:  "standard",
					CharFilter: []string{"amenity_filter"},
					Filter:     []string{"lowercase", "classic"},
				},
			},
			Normalizer: map[string]types.Normalizer{
				"custom_normalizer": &types.LowercaseNormalizer{
					Type: "lowercase",
				},
			},
			CharFilter: map[string]types.CharFilter{
				"amenity_filter": &types.PatternReplaceCharFilter{
					Type:        "pattern_replace",
					Pattern:     "\\|",
					Replacement: &emptyReplacement,
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
			// "categories": types.NestedProperty{
			// 	Type: "nested",
			// 	Properties: map[string]types.Property{
			// 		"id": types.KeywordProperty{
			// 			Type: "keyword",
			// 		},
			// 		"name": types.KeywordProperty{
			// 			Type:       "keyword",
			// 			Normalizer: &normalizer,
			// 		},
			// 	},
			// },

			"overallRate": types.FloatNumberProperty{
				Type: "float",
			},
			"wardId": types.KeywordProperty{
				Type: "keyword",
			},
			"lat": types.KeywordProperty{
				Type: "keyword",
			},
			"long": types.KeywordProperty{
				Type: "keyword",
			},
			"hostId": types.KeywordProperty{
				Type: "keyword",
			},

			"amenities": types.NestedProperty{
				Type: "nested",
				Properties: map[string]types.Property{
					"id": types.KeywordProperty{
						Type: "keyword",
					},
					"name": types.KeywordProperty{
						Type:       "keyword",
						Normalizer: &normalizer,
					},
				},
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
			"IsGuestFavor": types.BooleanProperty{
				Type: "boolean",
			},
			"title": types.TextProperty{
				Type:     "text",
				Analyzer: &customAnalyzer,
			},
			"body": types.TextProperty{
				Type: "text",
			},










			"owner": types.TextProperty{
				Type:     "text",
				Analyzer: &customAnalyzer,
			},
		},
	},
}
