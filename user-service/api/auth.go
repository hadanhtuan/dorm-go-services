package apiUser

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/model"
	"user-service/internal/model/enum"
	"user-service/internal/util"
	protoSdk "user-service/proto/sdk"
	protoUser "user-service/proto/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/common"
)

func (pc *UserAPI) Login(ctx context.Context, req *protoUser.MsgUser) (*protoSdk.BaseResponse, error) {
	filter := &model.User{}

	if req.Email == "" {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Email not null",
		})
	}

	filter.Email = req.Email

	result := model.UserDB.QueryOne(filter, nil)
	if result.Data == nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Username or password incorrect",
		})
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

func (pc *UserAPI) Register(ctx context.Context, req *protoUser.MsgUser) (*protoSdk.BaseResponse, error) {
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
		Email:    req.Email,
		Username: req.Username,
		Role:     &enum.UserRole.User,
		Password: hashPassword,
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

func (pc *UserAPI) UpdateUser(ctx context.Context, req *protoUser.MsgUser) (*protoSdk.BaseResponse, error) {
	user := &model.User{}
	db := model.UserDB.DB.Table(model.UserDB.TableName)

	if req.Email != "" {
		user.Email = req.Email
		db.Where("email = ?", req.Email)
	}

	if req.Username != "" {
		user.Username = req.Username
		db.Or("username = ?", req.Username)
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

	if req.Phone != "" {
		user.Phone = req.Phone
		db.Or("phone = ?", req.Phone)

	}

	var previousUser model.User
	db.First(&previousUser)

	if previousUser.ID != req.Id {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Error update user. Email/phone/username already exist",
		})
	}

	result := model.UserDB.Update(&model.User{ID: req.Id}, user)

	return util.ConvertToGRPC(result)

}

func (pc *UserAPI) RefreshToken(ctx context.Context, req *protoUser.MsgToken) (*protoSdk.BaseResponse, error) {
	result := model.LoginSessionDB.QueryOne(&model.LoginSession{
		RefreshToken: req.RefreshToken,
	}, nil)

	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Not found login session, please login again.",
		})
	}
	loginSession := result.Data.([]*model.LoginSession)[0]

	expireTime := loginSession.ExpiresAt

	//TODO: Can only refresh token only if refresh token not expires
	if expireTime < time.Now().Unix() {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Invalid refresh, please login again",
		})
	}

	token, err := CreateNewSeason(loginSession.UserId, loginSession.UserAgent, loginSession.IpAddress, loginSession.DeviceID)

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

func (pc *UserAPI) Logout(ctx context.Context, req *protoUser.MsgUser) (*protoSdk.BaseResponse, error) {

	filter := &model.LoginSession{
		UserId:   req.Id,
		DeviceID: req.DeviceId,
	}

	model.LoginSessionDB.Delete(filter)

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Logout successfully",
	})
}

func (pc *UserAPI) VerifyToken(ctx context.Context, req *protoUser.MsgToken) (*protoSdk.BaseResponse, error) {
	payload := common.JWTPayload{}

	jwt.ParseWithClaims(req.AccessToken, &payload, nil)

	result := model.LoginSessionDB.QueryOne(&model.LoginSession{
		UserId:   payload.UserID,
		DeviceID: payload.DeviceID,
	}, nil)

	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Not found login session, please login again.",
		})
	}

	loginSession := result.Data.([]*model.LoginSession)[0]

	jwtPayload, err := VerifyJWT(req.AccessToken, loginSession.SecretKey)

	if err != nil {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.Unauthorized,
			Message: "Error verify jwt. Error detail: " + err.Error(),
		})
	}
	return util.ConvertToGRPC(&common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   jwtPayload,
	})
}

func (pc *UserAPI) GetProfile(ctx context.Context, req *protoUser.MsgID) (*protoSdk.BaseResponse, error) {
	filter := map[string]interface{}{}
	filter["id"] = req.Id

	result := model.UserDB.QueryOne(filter, nil)
	return util.ConvertToGRPC(result)
}

func (pc *UserAPI) GetUsers(ctx context.Context, req *protoUser.MsgQueryUser) (*protoSdk.BaseResponse, error) {
	queryField := req.QueryFields
	filter := &model.User{}

	if queryField.Id != "" {
		filter.ID = queryField.Id
	}

	if queryField.Email != "" {
		filter.Email = queryField.Email
	}

	if queryField.Username != "" {
		filter.Username = queryField.Username
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

func (pc *UserAPI) GetUsersByIds(ctx context.Context, req *protoUser.MsgQueryUserByIds) (*protoSdk.BaseResponse, error) {

	result := model.UserDB.Query(req.Ids, req.Paginate.Offset, req.Paginate.Limit, nil)

	return util.ConvertToGRPC(result)
}

func CreateNewSeason(userID, userAgent, ipAddress, deviceID string) (*common.JWTToken, error) {
	expiresAt := time.Now().Add(3 * 24 * time.Hour) //TODO: token expire after 3 day
	refreshExpiresAt := expiresAt.Add(24 * time.Hour)

	// //TODO: each JWT have a unique login log ID
	// jwtPayload := &common.JWTPayload{
	// 	UserID:   userID,
	// 	DeviceID: deviceID,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(expiresAt),
	// 	},
	// }
	// token, err := aws.NewJWT(jwtPayload)

	jwtPayload := &common.JWTPayload{
		UserID:   userID,
		DeviceID: deviceID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	secretKey := sdk.RandomString(10)
	refreshToken := sdk.RandomString(25)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	accessToken, err := jwtToken.SignedString([]byte(secretKey))

	if err != nil {
		return nil, err
	}

	loginSession := &model.LoginSession{
		UserId:       userID,
		DeviceID:     deviceID,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		IpAddress:    ipAddress,
		SecretKey:    secretKey,
		ExpiresAt:    refreshExpiresAt.Unix(),
	}

	updater := &model.LoginSession{
		UserId:   userID,
		DeviceID: deviceID,
	}

	_ = model.LoginSessionDB.UpdateOrCreate(loginSession, updater)

	return &common.JWTToken{
		AccessToken:      accessToken,
		AccessExpiresAt:  expiresAt.Unix(),
		RefreshToken:     refreshToken,
		RefreshExpiresAt: refreshExpiresAt.Unix(),
	}, nil
}

func VerifyJWT(accessToken string, secret string) (*common.JWTPayload, error) {

	payload := common.JWTPayload{}

	claim, err := jwt.ParseWithClaims(accessToken, &payload, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !claim.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &payload, nil
}
