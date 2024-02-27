package apiUser

import (
	protoUser "user-service/proto/user"
)

type UserController struct {
	protoUser.UnimplementedUserServiceServer
}
