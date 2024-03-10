package api

import (
	"search-service/internal/amqp"
)

func (pc *SearchController) bindingMap() map[string]amqp.CallbackFunc {
	return map[string]amqp.CallbackFunc{
		amqp.PropertyCreated: pc.EventPropertyCreated,
	}
}

func (pc *SearchController) InitRoutingAMQP() {
	instance := amqp.GetConnection()
	bindingMap := pc.bindingMap()

	for routingKey, execFunc := range bindingMap {
		instance.BindingQueue(amqp.EXCHANGE, amqp.QUEUE, routingKey, execFunc)
	}

	instance.StartConsume(amqp.QUEUE)
}

func (pc *SearchController) EventPropertyCreated(data []byte) {

}
