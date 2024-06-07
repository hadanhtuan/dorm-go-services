package model

type AggregateResponse struct {
	DocCountErrorUpperBound int32             `json:"doc_count_error_upper_bound,omitempty"`
	SumOtherDocCount        int32             `json:"sum_other_doc_count,omitempty"`
	Buckets                 []AggregateBucket `json:"buckets,omitempty"`
}

type AggregateBucket struct {
	Key      string `json:"key,omitempty"`
	DocCount int32  `json:"doc_count,omitempty"`
}
