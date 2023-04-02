package sender

import (
	"time"

	"github.com/Shopify/sarama"
)

type Producer interface {
	SendMessage(key sarama.Encoder, value []byte) error
}

type MessageSender struct {
	syncProducer sarama.SyncProducer
	topic        string
}

func NewMessageSender(syncProducer sarama.SyncProducer, topic string) *MessageSender {
	return &MessageSender{syncProducer: syncProducer, topic: topic}
}

func (s *MessageSender) SendMessage(key sarama.Encoder, value []byte) error {
	_, _, err := s.syncProducer.SendMessage(&sarama.ProducerMessage{
		Key:       key,
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(value),
		Timestamp: time.Now(),
	})

	return err
}

func (s *MessageSender) Close() error {
	return s.syncProducer.Close()
}
