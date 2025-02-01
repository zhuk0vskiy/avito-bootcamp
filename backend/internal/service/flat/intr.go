package flat

import (
	"backend/internal/model/dto"
	"context"
)

type FlatIntr interface {
	Create(ctx context.Context, request *dto.CreateFlatRequest) (*dto.CreateFlatResponse, error)
	Delete(ctx context.Context, request *dto.CreateFlatResponse) error
	GetByID(ctx context.Context, request *dto.GetByIDRequest) (*dto.GetByIDResponse, error)
	GetByHouseID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
	UpdateStatus(ctx context.Context, request *dto.UpdateFlatStatusRequest) (*dto.UpdateFlatStatusResponse, error)
}