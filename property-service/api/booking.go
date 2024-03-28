package apiProperty

import (
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
	"context"

	"github.com/hadanhtuan/go-sdk/common"
)

// Property
func (bc *PropertyController) GetPropertyDetail(ctx context.Context, req *protoProperty.MsgGetProperty) (*protoSdk.BaseResponse, error) {

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Get Property Detail Successfully.",
	})
}
