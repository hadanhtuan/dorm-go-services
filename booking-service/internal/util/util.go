package util

import (
	protoSdk "booking-service/proto/sdk"
	"encoding/json"

	"github.com/hadanhtuan/go-sdk/common"
)

var (
	EXCHANGE = "searchExchange"
	QUEUE    = "searchQueue"

	// ROUTING KEY
	PropertyCreated = "property.property.created"
	PropertyUpdated = "property.property.updated"
)

func ConvertToGRPC(sdkResult *common.APIResponse) (*protoSdk.BaseResponse, error) {
	encodeData, _ := json.Marshal(sdkResult.Data)
	return &protoSdk.BaseResponse{
		Status:  sdkResult.Status,
		Message: sdkResult.Message,
		Data:    string(encodeData),
		Total:   sdkResult.Total,
	}, nil
}
