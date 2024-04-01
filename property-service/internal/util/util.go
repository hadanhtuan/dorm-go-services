package util

import (
	"encoding/json"
	protoSdk "property-service/proto/sdk"
	"reflect"

	"github.com/hadanhtuan/go-sdk/common"
)

var (
	SEARCH_EXCHANGE = "searchExchange"
	PROPERTY_EXCHANGE = "propertyExchange"
	PROPERTY_QUEUE    = "propertyQueue"
	SEARCH_QUEUE    = "searchQueue"

	// ROUTING KEY
	PaymentSuccess = "payment.success"
	PropertyCreated = "property.created"
	PropertyUpdated = "property.updated"
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

func ConvertEnumToSlice(obj interface{}) []string {
	result := []string{}

	v := reflect.ValueOf(obj)

	for i := 0; i < v.NumField(); i++ {
		result = append(result, v.Field(i).String())
	}

	return result
}
