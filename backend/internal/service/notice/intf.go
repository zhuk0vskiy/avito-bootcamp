package notice

import (
	"backend/internal/service/dto"
	"context"
	"errors"
)

// "backend/internal/service/dto"
// "context"

type NoticeIntf interface {
	Subscribe(ctx context.Context, request *dto.SubscribeRequest) (*dto.SubscribeResponse, error)
	// Unsubscribe(ctx context.Context, request *dto.UnsubscribeRequest) (*dto.UnsubscribeResponse, error)
	IsNeedToNoticeSubscribers(ctx context.Context, request *dto.IsNeedToNoticeSubscribersRequest) ()
	SendNotices(ctx context.Context, request *dto.SendNoticesRequest) (*dto.SendNoticesResponse, error)
	// Noticeing(ctx co)
	NoticeSubscribers(ctx context.Context, request *dt)
}

var (
	ErrNilRequest = errors.New("dto request is nil")
	ErrNilContext = errors.New("context is nil")
	ErrBadHouseID      = errors.New("bad house id")
	ErrBadUserID = errors.New("bad user id")
	ErrBadApartmentID = errors.New("bad apartment id")
)
