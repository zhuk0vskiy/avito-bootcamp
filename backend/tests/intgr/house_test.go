package intgr

import (
	"backend/config"
	"backend/internal/repo/postgres"
	houseRepo "backend/internal/repo/postgres/house"
	serviceDto "backend/internal/service/dto"
	houseService "backend/internal/service/house"
	loggerMock "backend/pkg/logger/mocks"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type HouseSuite struct {
	suite.Suite
}

func (suite *HouseSuite) TestHouseAddSuccess_01(t provider.T) {
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

	houseR := houseRepo.NewHouseRepo(mockLogger, r)
	houseS := houseService.NewHouseService(mockLogger, houseR)

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Times(0)
	mockLogger.On("Infof", mock.Anything, mock.Anything).Times(0)
	mockLogger.On("Warnf", mock.Anything, mock.Anything).Times(0)

	id, err := uuid.Parse("54930a14-ffab-46e9-ac9f-7f9c50341872")
	if err != nil {
		fmt.Println("failed to parse")
		return
	}
	response, err := houseS.Create(ctx, &serviceDto.CreateHouseRequest{
		CreatorID:     id,
		Address:       "Moscow, krasna square",
		MaxApartments: 4,
	})
	fmt.Println("----")

	assert.NoError(t, err)
	assert.NotEmpty(t, response)

}
