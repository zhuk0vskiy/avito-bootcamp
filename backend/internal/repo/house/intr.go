package house

import (
	"backend/internal/repo/dto"
	"context"
)

type HouseIntr interface {
	Add(ctx context.Context, request *dto.AddHouseRequest) (*dto.AddHouseResponse, error)
	GetByID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
	

}