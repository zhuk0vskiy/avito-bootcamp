package notice

import (
	repoDto "backend/internal/repo/dto"
	"context"
	"errors"
)

// "backend/internal/repo/dto"
// "context"

type NoticeIntf interface {
	Subscribe(ctx context.Context, request *repoDto.SubscribeRequest) (*repoDto.SubscribeResponse, error)
	GetSubscribersByHouseID(ctx context.Context, request *repoDto.GetSubscribersByHouseIDRequest) (*repoDto.GetSubscribersByHouseIDResponse, error)
	// Unsubscribe(ctx context.Context, request *dto.UnsubscribeRequest) (*dto.UnsubscribeResponse, error)
	// Noticeing(ctx context.Context, request *)
	// SendNotice(ctx context.Context, request *)
	GetNoticesOutbox(ctx context.Context, request *repoDto.GetNoticesOutboxRequest) (*repoDto.GetNoticesOutboxResponse, error)
	Add(ctx context.Context, request *repoDto.AddNoticeRequest) (*repoDto.AddNoticeResponse, error)
	ConfirmNoticeOutbox(ctx context.Context, request *repoDto.ConfirmNoticeOutboxRequest) (*repoDto.ConfirmNoticeOutboxResponse, error)
}

var (
	ErrNilRequest = errors.New("noice add dto request is nil")
	ErrNilContext = errors.New("notice add context is nil")
	ErrStartTrans = errors.New("failed to init transaction")
	ErrWhileTrns = errors.New("failed while transaction -- doing rollback")
	ErrRollback = errors.New("failed rollback")
	ErrAddNotice = errors.New("failed to add notice")
	ErrAddNoticeOutbox = errors.New("failed to add notice into outbox")
	ErrCommit = errors.New("failed to commit")
)

var (
	ErrExec       = errors.New("error while exec")
	ErrQueryRow = errors.New("error while quering row")
	ErrQuery      = errors.New("error while query")
)



