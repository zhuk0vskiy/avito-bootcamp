package apartment

import (
	"backend/internal/service/dto"
	"context"
)

type apartmentIntf interface {
	Create(ctx context.Context, request *dto.CreateapartmentRequest) (*dto.CreateapartmentResponse, error)
	Delete(ctx context.Context, request *dto.CreateapartmentResponse) error
	GetByID(ctx context.Context, request *dto.GetByIDRequest) (*dto.GetByIDResponse, error)
	GetByHouseID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
	UpdateStatus(ctx context.Context, request *dto.UpdateapartmentStatusRequest) (*dto.UpdateapartmentStatusResponse, error)
}
