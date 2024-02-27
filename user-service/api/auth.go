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

	filter := map[string]interface{}{}

	filter["email"] = req.Email

	result := model.UserDB.QueryOne(filter)
	if result.Data == nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Username or password incorrect",
		}, nil
	}
	data := result.Data.([]*model.User)[0]

	isVerify := sdk.VerifyPassword(req.Password, data.Password)
	if !isVerify || result.Data == nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Username or password incorrect",
		}, nil
	}

	jwtPayload := &common.JWTPayload{
		ID:    data.ID,
		Email: data.Email,
		DeviceID: req.DeviceId,
	}
	token, err := aws.NewJWT(jwtPayload)

	if err != nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error generate JWT. Error Detail: " + err.Error(),
		}, nil
	}

	//TODO: save login log
	loginLog := &model.LoginLog{
		UserId:    data.ID,
		UserAgent: req.UserAgent,
		IpAddress: req.IpAddress,
		DeviceID:  req.DeviceId,
	}
	model.LoginLogDB.Create(loginLog)

	encodeData, _ := json.Marshal(token)
	return &protoSdk.BaseResponse{
		Status:  result.Status,
		Message: result.Message,
		Data:    string(encodeData),
		Total:   result.Total,
	}, nil
}

func (pc *UserController) Register(ctx context.Context, req *protoUser.MsgRegister) (*protoSdk.BaseResponse, error) {
	filter := map[string]interface{}{}

	filter["email"] = req.Email
	checkExist := model.UserDB.QueryOne(filter)

	if checkExist.Data != nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Email already exist",
		}, nil
	}

	hashPassword, _ := sdk.HashPassword(req.Password)
	user := &model.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      &enum.UserRole.User,
		Password:  hashPassword,
	}

	result := model.UserDB.Create(user)
	data := result.Data.([]*model.User)[0]

	jwtPayload := &common.JWTPayload{
		ID:    data.ID,
		Email: data.Email,
		DeviceID: req.DeviceId,
	}

	token, err := aws.NewJWT(jwtPayload)

	if err != nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error generate JWT. Error Detail: " + err.Error(),
		}, nil
	}

	//TODO: save login log
	loginLog := &model.LoginLog{
		UserId:    data.ID,
		UserAgent: req.UserAgent,
		IpAddress: req.IpAddress,
		DeviceID:  req.DeviceId,
	}
	model.LoginLogDB.Create(loginLog)

	encodeData, _ := json.Marshal(token)
	return &protoSdk.BaseResponse{
		Status:  result.Status,
		Message: result.Message,
		Data:    string(encodeData),
		Total:   result.Total,
	}, nil
}
