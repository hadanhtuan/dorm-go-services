package apiProperty

import (
	protoProperty "property-service/proto/property"
	protoUser "property-service/proto/user"
)

type PropertyAPI struct {
	protoProperty.UnimplementedPropertyServiceServer

	//client
	UserServiceClient protoUser.UserServiceClient
}

func InitAPI(userClient protoUser.UserServiceClient) {
	pa := &PropertyAPI{
		UserServiceClient: userClient,
	}
	pa.InitRoutingAMQP()

	InstanceAPI = pa
}

var (
	InstanceAPI *PropertyAPI
)
