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

func InitAPI() {
	pa := &PropertyAPI{}
	pa.InitRoutingAMQP()

	InstanceAPI = pa
}

var (
	InstanceAPI *PropertyAPI
)
