package util

import (
	protoSdk "booking-service/proto/sdk"
	"encoding/json"

	"github.com/hadanhtuan/go-sdk/common"
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
