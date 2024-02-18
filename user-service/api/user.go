package apiUser

import (
	"context"
	"encoding/json"
	"user-service/internal/model"
	protoSdk "user-service/proto/sdk"
	protoUser "user-service/proto/user"
)

type UserController struct {
	protoUser.UnimplementedUserServiceServer
}

func (pc *UserController) Login(ctx context.Context, req *protoUser.MessageLogin) (*protoSdk.BaseResponse, error) {
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	result := model.UserDB.Create(user)
	// result := model.UserDB.Query(user, 1, 10)

	data := result.Data.([]*model.User)

	encodeData, _ := json.Marshal(data)
	baseResponse := &protoSdk.BaseResponse{
		Status:  result.Status,
		Message: result.Message,
		Data:    string(encodeData),
		Total:   result.Total,
	}

	return baseResponse, nil
}