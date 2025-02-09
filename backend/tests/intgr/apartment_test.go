package intgr

import (
	"backend/config"
	"backend/internal/repo/postgres"
	// "backend/internal/repo/postgres/apartment"
	apartmentRepo "backend/internal/repo/postgres/apartment"
	serviceDto "backend/internal/service/dto"
	apartmentService "backend/internal/service/apartment"
	loggerMock "backend/pkg/logger/mocks"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ApartmentSuite struct {
	suite.Suite
}

func (suite *ApartmentSuite) TestApartmentAddSuccess_01(t provider.T) {
	mockLogger := new(loggerMock.Interface)
	ctx := context.Background()
	cfg := config.Postgres{
		Host:     "localhost",
		Port:     5438,
		User:     "avito",
		Password: "avito",
		Database: "avito",
		Driver:   "postgres",
	}
	dbConnector, err := postgres.NewDbConn(ctx, &cfg)
	if err != nil {
		fmt.Println("db is not connected")
		return
	}
	r := postgres.NewRetryAdapter(dbConnector, 2, 2)

	apartmentR := apartmentRepo.NewApartmentRepo(mockLogger, r)
	houseS := apartmentService.NewApartmentService(mockLogger, apartmentR)

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Times(0)
	mockLogger.On("Infof", mock.Anything, mock.Anything).Times(0)
	mockLogger.On("Warnf", mock.Anything, mock.Anything).Times(0)

	id, err := uuid.Parse("54930a14-ffab-46e9-ac9f-7f9c50341872")
	if err != nil {
		fmt.Println("failed to parse")
		return
	}

	houseID,err := uuid.Parse("b5297fe6-a5d9-43a8-9923-b73daabc8ceb")
	if err != nil {
		fmt.Println("failed to parse")
		return
	}
	response, err := houseS.Create(ctx, &serviceDto.CreateApartmentRequest{
		CreatorID:     id,
		HouseID: houseID,
		Price: 4000000,
		Rooms: 4,
	})
	fmt.Println("----")

	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}