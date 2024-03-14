package api

import (
	"encoding/json"
	"fmt"
	"search-service/internal/model"
	"search-service/internal/util"

	"github.com/hadanhtuan/go-sdk/amqp"
)

func (pc *SearchController) bindingMap() map[string]amqp.CallbackFunc {
	return map[string]amqp.CallbackFunc{
		util.PropertyCreated: pc.EventPropertyCreated,
	}
}

func (pc *SearchController) InitRoutingAMQP() {
	instance := amqp.GetConnection()
	bindingMap := pc.bindingMap()

	for routingKey, execFunc := range bindingMap {
		instance.BindingQueue(util.EXCHANGE, util.QUEUE, routingKey, execFunc)
	}

	instance.StartConsume(util.QUEUE)
}

func (pc *SearchController) EventPropertyCreated(payload []byte) {
	var data model.Property
	json.Unmarshal(payload, &data)
	fmt.Println(data.ID)
}
