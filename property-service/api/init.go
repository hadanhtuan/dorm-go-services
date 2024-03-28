package apiProperty

import (
	protoProperty "property-service/proto/property"
)

type PropertyController struct {
	protoProperty.UnimplementedPropertyServiceServer
}
