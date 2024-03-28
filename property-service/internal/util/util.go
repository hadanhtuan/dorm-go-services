package util

import (
	protoSdk "property-service/proto/sdk"
	"encoding/json"

	"github.com/hadanhtuan/go-sdk/common"
)

var (
	EXCHANGE = "searchExchange"
	QUEUE    = "searchQueue"

	// ROUTING KEY
	PropertyCreated = "property.property.created"
	PropertyUpdated = "property.property.updated"
	PropertyDeleted = "property.property.deleted"
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

func ConvertStruct[T any](obj any) T {
	var mergeObj T

	byteObj, _ := json.Marshal(obj)

	json.Unmarshal(byteObj, &mergeObj)

	return mergeObj
}

func ConvertSlice[T any, K any](slice []K) []T {
	var result []T

	for _, obj := range slice {
		item := ConvertStruct[T](obj)
		result = append(result, item)
	}

	return result
}
