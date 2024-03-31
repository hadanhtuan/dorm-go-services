package model

type CountByStatus struct {
	Status   string `json:"status,omitempty"`
	Quantity *int64  `json:"quantity,omitempty"`
}
