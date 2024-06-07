package model

type SearchTrackingDocument struct {
	SearchCount int64  `json:"searchCount"`
	SearchText  string `json:"searchText,omitempty"`
	Highlight   string `json:"highlight,omitempty"`
}
