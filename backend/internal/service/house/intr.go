package house

import (
	"backend/internal/service/dto"
	"context"
)

type HouseIntr interface {
	Create(ctx context.Context, request *dto.CreateHouseRequest) (*dto.CreateHouseResponse, error)
	Delete(ctx context.Context, request *dto.DeleteHouseRequest) error
	GetByID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
}