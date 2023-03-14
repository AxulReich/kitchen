package sender

import (
	"time"

	"github.com/Shopify/sarama"
)

type Producer interface {
	SendMessage(value []byte) error
}

type MessageSender struct {
	syncProducer sarama.SyncProducer
	topic        string
}

func NewMessageSender(syncProducer sarama.SyncProducer, topic string) *MessageSender {
	return &MessageSender{syncProducer: syncProducer, topic: topic}
}

func (s *MessageSender) SendMessage(value []byte) error {
	// By default the package uses HashPartitioner
	// Messages with the same Key will be added to the single partition
	// In case Key is omited — the package uses RandomPartitioner
	_, _, err := s.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(value),
		Timestamp: time.Now(),
	})

	return err
}

func (s *MessageSender) Close() {
	s.syncProducer.Close()
}
