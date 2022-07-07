package service_producer

import (
	"encoding/json"
	"reflect"

	"github.com/Shopify/sarama"
)


type EventProducer interface {
	Produce(event interface{}) error 
}

type eventProducer struct {
	producer sarama.SyncProducer
} 

func NewEventProducer  (producer sarama.SyncProducer) EventProducer {
	 return eventProducer{producer}
}


func (p eventProducer) Produce(event interface{}) error {
	topic := reflect.TypeOf(event).Name()
  
	value , err := json.Marshal(event)   
	if err != nil {
		return err 
	}

	msg := sarama.ProducerMessage{
		   Topic: topic,
		   Value: sarama.ByteEncoder(value),
	}
	_,_,err = p.producer.SendMessage(&msg)
	if err != nil {
		return err 
	}  
	return nil 
}