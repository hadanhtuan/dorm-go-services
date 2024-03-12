package es

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)
//constant do not have memory address(cannot use &)
var (
	emptyReplacement = " "
	analyzer         = "custom_analyzer"
	normalizer       = "custom_normalizer"

	CacheSearchDocument = "search:document"
	CacheSearchTracking = "search:tracking"

	SearchAll  = "ALL"
	SearchPost = "POST"
	SearchUser = "USER"

	PropertyIndices           = "property_index"
	UserIndices           = "user_index"
	SearchTrackingIndices = "search_tracking_index"
)

var PropertyIndicesCnf = &create.Request{
	Settings: &types.IndexSettings{
		Analysis: &types.IndexSettingsAnalysis{
			Analyzer: map[string]types.Analyzer{
				"custom_analyzer": &types.CustomAnalyzer{
					Tokenizer:  "standard",
					CharFilter: []string{"tag_filter"},
					Filter:     []string{"lowercase", "classic"},
				},
			},
			Normalizer: map[string]types.Normalizer{
				"custom_normalizer": &types.LowercaseNormalizer{
					Type: "lowercase",
				},
			},
			CharFilter: map[string]types.CharFilter{
				"tag_filter": &types.PatternReplaceCharFilter{
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
			"cityId": types.KeywordProperty{
				Type: "keyword",
			},
			"countryId": types.KeywordProperty{
				Type: "keyword",
			},
			"categories": types.NestedProperty{
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
			"tags": types.NestedProperty{
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
			"title": types.TextProperty{
				Type:     "text",
				Analyzer: &analyzer,
			},
			"content": types.TextProperty{
				Type: "text",
			},
			"location": types.TextProperty{
				Type: "text",
			},
			"isPublic": types.BooleanProperty{
				Type: "boolean",
			},
			"viewCount": types.IntegerNumberProperty{
				Type: "integer",
			},
			"readCount": types.IntegerNumberProperty{
				Type: "integer",
			},
			"owner": types.TextProperty{
				Type:     "text",
				Analyzer: &analyzer,
			},
			"createdAt": types.DateProperty{
				Type: "date",
			},
		},
	},
}
