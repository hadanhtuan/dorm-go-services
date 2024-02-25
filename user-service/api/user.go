package apiUser

import (
	"context"
	"encoding/json"
	"user-service/internal/model"
	"user-service/internal/model/enum"
	protoSdk "user-service/proto/sdk"
	protoUser "user-service/proto/user"
)

func (pc *UserController) Login(ctx context.Context, req *protoUser.MsgLogin) (*protoSdk.BaseResponse, error) {
	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	result := model.UserDB.Create(user)

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

func (pc *UserController) Register(ctx context.Context, req *protoUser.MsgRegister) (*protoSdk.BaseResponse, error) {
	user := &model.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      &enum.UserRole.User,
	}

	result := model.UserDB.Create(user)

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
