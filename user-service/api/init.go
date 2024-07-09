package apiUser

import (
	protoUser "user-service/proto/user"
)

type UserAPI struct {
	protoUser.UnimplementedUserServiceServer
}

func InitAPI() {
	ua := &UserAPI{
	}

	InstanceAPI = ua
}


var (
	InstanceAPI *UserAPI
)
