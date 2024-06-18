package util

import (
	"encoding/json"
	protoSdk "search-service/proto/sdk"

	"github.com/hadanhtuan/go-sdk/common"
)

var (
	SEARCH_EXCHANGE = "searchExchange"
	SEARCH_QUEUE    = "searchQueue"

	// ROUTING KEY
	PropertyCreated = "property.created"
	PropertyUpdated = "property.updated"

	Desc = "desc"
	Asc  = "asc"
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

func UniqueSliceElements[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]bool, len(inputSlice))
	for _, element := range inputSlice {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}
	return uniqueSlice
}
