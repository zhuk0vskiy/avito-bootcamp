package notify

import (
	"backend/internal/service/dto"
	"context"
)

type NotifyIntf interface {
	Subscribe(ctx context.Context, request *dto.SubscribeRequest) (*dto.SubscribeResponse, error)
	Unsubscribe(ctx context.Context, request *dto.UnsubscribeRequest) (*dto.UnsubscribeResponse, error)
	// Notifying(ctx co)
}
