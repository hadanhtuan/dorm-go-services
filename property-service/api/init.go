package apiProperty

import (
	protoProperty "property-service/proto/property"
)

type PropertyController struct {
	protoProperty.UnimplementedPropertyServiceServer
}

func (pc *PropertyController) InitController() {
	pc.InitRoutingAMQP()
}
