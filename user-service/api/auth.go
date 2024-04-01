package apiUser

import (
	"context"
	"time"
	"user-service/internal/model"
	"user-service/internal/model/enum"
	"user-service/internal/util"
	protoSdk "user-service/proto/sdk"
	protoUser "user-service/proto/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/aws"
	"github.com/hadanhtuan/go-sdk/common"
)

func (pc *UserController) Login(ctx context.Context, req *protoUser.MsgUser) (*protoSdk.BaseResponse, error) {
	filter := &model.User{}

	filter.Email = req.Email

	result := model.UserDB.QueryOne(filter, nil)
	if result.Data == nil {
		return &protoSdk.BaseResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Username or password incorrect",
		}, nil
	}
	data := result.Data.([]*model.User)[0]

	isVerify := sdk.VerifyPassword(req.Password, data.Password)
	if !isVerify || result.Data == nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Username or password incorrect",
		})
	}
	token, err := CreateNewSeason(data.ID, req.UserAgent, req.IpAddress, req.DeviceId)

	if err != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Ok,
			Message: "Error generate JWT. Error Detail: " + err.Error(),
		})
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   token,
	})
}

func (pc *UserController) Register(ctx context.Context, req *protoUser.MsgUser) (*protoSdk.BaseResponse, error) {
	filter := &model.User{}

	filter.Email = req.Email
	checkExist := model.UserDB.QueryOne(filter, nil)

	if checkExist.Data != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Email already exist",
		})
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

	token, err := CreateNewSeason(data.ID, req.UserAgent, req.IpAddress, req.DeviceId)

	if err != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Ok,
			Message: "Error generate JWT. Error Detail: " + err.Error(),
		})
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   token,
	})
}

func (pc *UserController) UpdateUser(ctx context.Context, req *protoUser.MsgUser) (*protoSdk.BaseResponse, error) {
	user := &model.User{}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.IsActive != nil {
		user.IsActive = req.IsActive
	}

	if req.Gender != "" {
		user.Gender = req.Gender
	}
	
	if req.Password != "" {
		hashPassword, _ := sdk.HashPassword(req.Password)
		user.Password = hashPassword
	}

	result := model.UserDB.Update(&model.User{ID: req.Id}, user)

	return util.ConvertToGRPC(result)

}

func (pc *UserController) RefreshToken(ctx context.Context, req *protoUser.MsgToken) (*protoSdk.BaseResponse, error) {
	jwtPayload, err := aws.VerifyJWT(req.Token)

	if err != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Error verify jwt. Error detail: " + err.Error(),
		})
	}
	expireTime := jwtPayload.RegisteredClaims.ExpiresAt.Time
	deadlineOutdate := time.Now().Add(24 * time.Hour)

	//TODO: Can only refresh token if token expires within 1 day
	if expireTime.Before(time.Now()) && expireTime.After(deadlineOutdate) {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Invalid refresh, please login again",
		})
	}

	filter := map[string]interface{}{}
	filter["user_id"] = jwtPayload.UserID
	filter["id"] = jwtPayload.LoginLogID
	filter["device_id"] = jwtPayload.DeviceID
	// filter["expires_at"] = expireTime
	filter["is_logout"] = false

	result := model.LoginLogDB.QueryOne(filter, nil)

	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Not found season, please login again",
		})
	}

	loginLog := result.Data.([]*model.LoginLog)[0]

	if loginLog.ExpiresAt.Before(deadlineOutdate) || loginLog.ExpiresAt.Before(expireTime) {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Season outdate, please login again",
		})
	}

	loginLog.IsLogout = true
	model.LoginLogDB.Update(filter, loginLog) //TODO: delete old season

	token, err := CreateNewSeason(jwtPayload.UserID, loginLog.UserAgent, loginLog.IpAddress, jwtPayload.DeviceID)

	if err != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Ok,
			Message: "Error generate JWT. Error Detail: " + err.Error(),
		})
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Refresh token successfully",
		Data:    token,
	})
}

func (pc *UserController) Logout(ctx context.Context, req *protoUser.MsgID) (*protoSdk.BaseResponse, error) {

	filter := map[string]interface{}{}
	filter["id"] = req.Id
	filter["is_logout"] = false

	loginLog := &model.LoginLog{
		ExpiresAt: time.Now(),
		IsLogout:  true,
	}

	model.LoginLogDB.Update(filter, loginLog)

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Logout successfully",
	})
}

func (pc *UserController) VerifyToken(ctx context.Context, req *protoUser.MsgToken) (*protoSdk.BaseResponse, error) {
	jwtPayload, _ := aws.VerifyJWT(req.Token)
	expireTime := jwtPayload.RegisteredClaims.ExpiresAt.Time

	if expireTime.Before(time.Now()) {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Token timeout, please login again",
		})
	}

	filter := map[string]interface{}{}
	filter["user_id"] = jwtPayload.UserID
	filter["id"] = jwtPayload.LoginLogID
	filter["device_id"] = jwtPayload.DeviceID
	filter["is_logout"] = false

	result := model.LoginLogDB.QueryOne(filter, nil)

	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Not found season, please login again",
		})
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   jwtPayload,
	})
}

func (pc *UserController) GetProfile(ctx context.Context, req *protoUser.MsgID) (*protoSdk.BaseResponse, error) {
	filter := map[string]interface{}{}
	filter["id"] = req.Id

	result := model.UserDB.QueryOne(filter, nil)
	return util.ConvertToGRPC(result)
}

func (pc *UserController) GetUsers(ctx context.Context, req *protoUser.MsgQueryUser) (*protoSdk.BaseResponse, error) {
	queryField := req.QueryFields
	filter := &model.User{}

	if queryField.Id != "" {
		filter.ID = queryField.Id
	}

	if queryField.Email != "" {
		filter.Email = queryField.Email
	}

	if queryField.Gender != "" {
		filter.Gender = queryField.Gender
	}

	if queryField.Role != nil {
		role := enum.UserRoleValue(*queryField.Role)
		filter.Role = &role
	}

	if queryField.Phone != "" {
		filter.Phone = queryField.Phone
	}

	if queryField.IsActive != nil {
		filter.IsActive = queryField.IsActive
	}

	result := model.UserDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, nil)
	return util.ConvertToGRPC(result)
}

func CreateNewSeason(userID, userAgent, ipAddress, deviceID string) (*common.JWTToken, error) {
	expiresAt := time.Now().Add(3 * 24 * time.Hour) //TODO: token expire after 3 day

	result := model.LoginLogDB.Create(&model.LoginLog{ //TODO: create new season
		UserId:    userID,
		UserAgent: userAgent,
		IpAddress: ipAddress,
		DeviceID:  deviceID,
		ExpiresAt: expiresAt,
		IsLogout:  false,
	})
	loginLog := result.Data.([]*model.LoginLog)[0]

	//TODO: each JWT have a unique login log ID
	jwtPayload := &common.JWTPayload{
		UserID:     userID,
		LoginLogID: loginLog.ID,
		DeviceID:   deviceID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token, err := aws.NewJWT(jwtPayload)

	return token, err
}
