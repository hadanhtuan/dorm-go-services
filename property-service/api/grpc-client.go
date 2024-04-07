package apiProperty

import (
	"context"
	"property-service/internal/util"
	"time"

	protoSdk "property-service/proto/sdk"
	protoUser "property-service/proto/user"

	"github.com/hadanhtuan/go-sdk/common"
)

func (pc *PropertyController) GetUsers(listIds []string) *common.APIResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	payload := &protoUser.MsgQueryUserByIds{
		Paginate: &protoSdk.Pagination{
			Offset: 0,
			Limit:  int32(len(listIds)),
		},
		Ids: listIds,
	}

	result, _ := pc.UserServiceClient.GetUsersByIds(ctx, payload)
	return util.ConvertGrpcResult(result)
}
