package apiUser

import (
	"context"
	"encoding/json"
	"user-service/internal/model"
	"user-service/internal/model/enum"
	protoSdk "user-service/proto/sdk"
	protoUser "user-service/proto/user"

	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/aws"
	"github.com/hadanhtuan/go-sdk/common"
)

func (pc *UserController) Login(ctx context.Context, req *protoUser.MsgLogin) (*protoSdk.BaseResponse, error) {
	user := &model.User{
		Email: req.Email,
	}

	result := model.UserDB.QueryOne(user)

	data := result.Data.([]*model.User)[0]

	isVerify := sdk.VerifyPassword(req.Password, data.Password)
	if !isVerify || result.Data == nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Username or password incorrect",
		}, nil
	}

	jwtPayload := &common.JWTPayload{
		ID:       data.ID,
		Username: data.Username,
	}

	token, err := aws.NewJWT(jwtPayload)

	if err != nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error generate JWT. Error Detail: " + err.Error(),
		}, nil
	}

	encodeData, _ := json.Marshal(map[string]interface{}{
		"token": token,
	})

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
		Email: req.Email,
	}
	checkExist := model.UserDB.QueryOne(user)

	if checkExist.Data != nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Email already exist",
		}, nil
	}

	hashPassword, _ := sdk.HashPassword(req.Password)

	user.Password = hashPassword
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Role = &enum.UserRole.User

	result := model.UserDB.Create(user)

	data := result.Data.([]*model.User)[0]
	jwtPayload := &common.JWTPayload{
		ID:       data.ID,
		Username: data.Username,
	}

	token, err := aws.NewJWT(jwtPayload)

	if err != nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error generate JWT. Error Detail: " + err.Error(),
		}, nil
	}

	encodeData, _ := json.Marshal(map[string]interface{}{
		"token": token,
	})

	baseResponse := &protoSdk.BaseResponse{
		Status:  result.Status,
		Message: result.Message,
		Data:    string(encodeData),
		Total:   result.Total,
	}

	return baseResponse, nil
}
