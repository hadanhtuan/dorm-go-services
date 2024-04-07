package apiProperty

import (
	protoProperty "property-service/proto/property"
	protoUser "property-service/proto/user"
)

type PropertyController struct {
	protoProperty.UnimplementedPropertyServiceServer

	//client
	UserServiceClient protoUser.UserServiceClient
}

func (pc *PropertyController) InitController() {
	pc.InitRoutingAMQP()
}
