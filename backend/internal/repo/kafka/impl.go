package kafka

import (
	"backend/pkg/logger"
	"context"

	"encoding/json"

	repoDto "backend/internal/repo/dto"

	"github.com/IBM/sarama"
)

type Producer struct {
	logger   logger.Interface
	producer sarama.SyncProducer
	topic    string
}

func NewProducer(logger logger.Interface, brokers []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, ErrProducerCreate
	}

	return &Producer{
		logger:   logger,
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}

func (p *Producer) ProduceMessages(ctx context.Context, request *repoDto.ProduceMessageRequest) (*repoDto.ProduceMessageResponse, error) {
	method := "KafkaRepo -- ProduceMessage"
	messages := make([]*sarama.ProducerMessage, 0)

	for i := 0; i < len(request.Messages); i++ {
		jsonData, err := json.Marshal(request.Messages[i])
		if err != nil {
			p.logger.Errorf("%s -- %s -- %s", method, ErrMarshalResponse, err)
			return nil, ErrMarshalResponse
		}

		message := &sarama.ProducerMessage{
			Key:   sarama.StringEncoder("test-key"),
			Topic: p.topic,
			Value: sarama.StringEncoder(jsonData),
			Partition: 1,

		}
		messages = append(messages, message)
	}

	err := p.producer.SendMessages(messages)
	if err != nil {
		p.logger.Errorf("%s -- %s -- %s", method, ErrProduce, err)
		return nil, ErrProduce
	}

	// p.producer.SendMessage()

	p.logger.Infof("Message sent to topic -- \n", p.topic)
	return nil, nil
}
