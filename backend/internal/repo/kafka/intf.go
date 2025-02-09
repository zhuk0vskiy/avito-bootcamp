package kafka

import (
	repoDto "backend/internal/repo/dto"
	"context"
	"errors"
)

type KafkaIntf interface {
	ProduceMessages(ctx context.Context, request *repoDto.ProduceMessageRequest) (*repoDto.ProduceMessageResponse, error)
}

var (
	ErrMarshalResponse = errors.New("failed to marshal notice model from response data")
	ErrProduce = errors.New("failed to produce msg into topic")
	ErrProducerCreate = errors.New("failed to create producer")
)
