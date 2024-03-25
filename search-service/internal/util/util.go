package util

import (
	"encoding/json"
	protoSdk "search-service/proto/sdk"

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


func MergeStruct(target, obj any) []byte {
	var mergeObj map[string]any

	byteTarget, _ := json.Marshal(target)
	byteObj, _ := json.Marshal(obj)

	json.Unmarshal(byteTarget, &mergeObj)
	json.Unmarshal(byteObj, &mergeObj)

	byteMerged, _ := json.Marshal(mergeObj)

	return byteMerged
}