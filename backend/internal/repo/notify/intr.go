package notify

import (
	"backend/internal/repo/dto"
	"context"
)

type NotifyIntr interface {
	Subscribe(ctx context.Context, request *dto.SubscribeRequest) (*dto.SubscribeResponse, error)
	Unsubscribe(ctx context.Context, request *dto.UnsubscribeRequest) (*dto.UnsubscribeResponse, error)
	// Notifying(ctx context.Context, request *)
}