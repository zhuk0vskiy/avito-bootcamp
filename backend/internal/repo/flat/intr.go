package flat

import (
	"backend/internal/repo/dto"
	"context"
)

type FlatIntr interface {
	Add(ctx context.Context, request *dto.AddFlatRequest) (*dto.AddFlatResponse, error)
	GetByID(ctx context.Context, request *dto.GetByIDRequest) (*dto.GetByIDResponse, error)
	GetByHouseID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
	UpdateStatus(ctx context.Context, request *dto.UpdateFlatStatusRequest) (*dto.UpdateFlatStatusResponse, error)
}