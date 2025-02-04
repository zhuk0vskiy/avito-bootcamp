package apartment

import (
	"backend/internal/repo/dto"
	"context"
)

type ApartmentIntf interface {
	Add(ctx context.Context, request *dto.AddApartmentRequest) (*dto.AddApartmentResponse, error)
	GetByID(ctx context.Context, request *dto.GetByIDRequest) (*dto.GetByIDResponse, error)
	GetByHouseID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
	UpdateStatus(ctx context.Context, request *dto.UpdateapartmentStatusRequest) (*dto.UpdateapartmentStatusResponse, error)
}
