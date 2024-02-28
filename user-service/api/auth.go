package apiUser

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
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
		ID:       data.ID,
		Email:    data.Email,
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
		ExpiresAt: jwtPayload.ExpiresAt.Time,
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
	fmt.Println("motherfucker")

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
		ID:       data.ID,
		Email:    data.Email,
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
		ExpiresAt: jwtPayload.ExpiresAt.Time,
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

func (pc *UserController) RefreshToken(ctx context.Context, req *protoUser.MsgToken) (*protoSdk.BaseResponse, error) {
	jwtPayload, _ := aws.VerifyJWT(req.Token)
	expireTime := jwtPayload.RegisteredClaims.ExpiresAt.Time

	if time.Duration(expireTime.Hour()) > 24*time.Hour {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Token outdate, please login",
		}, nil
	}

	filter := map[string]interface{}{}
	filter["expires_at <"] = expireTime
	filter["email"] = req.Email
	filter["device_id"] = req.DeviceId

	result := model.LoginLogDB.QueryOne(filter)

	if result.Status == common.APIStatus.NotFound {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Not found season, please login",
		}, nil
	}

	token, _ := aws.NewJWT(jwtPayload)
	loginLog := result.Data.([]*model.LoginLog)[0]
	loginLog.ExpiresAt = token.ExpiresAt

	model.LoginLogDB.Update(filter, loginLog)

	encodeData, _ := json.Marshal(token)
	return &protoSdk.BaseResponse{
		Status:  common.APIStatus.Ok,
		Message: "Refresh token successfully",
		Data:    string(encodeData),
	}, nil
}

func (pc *UserController) Logout(ctx context.Context, req *protoUser.MsgToken) (*protoSdk.BaseResponse, error) {
	filter := map[string]interface{}{}
	filter["email"] = req.Email
	filter["device_id"] = req.DeviceId

	loginLog := &model.LoginLog{
		ExpiresAt: time.Now(),
	}

	model.LoginLogDB.Update(filter, loginLog)

	return &protoSdk.BaseResponse{
		Status:  common.APIStatus.Ok,
		Message: "Logout successfully",
	}, nil
}
