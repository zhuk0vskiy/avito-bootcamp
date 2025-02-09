package intgr

import (
	"backend/config"
	"backend/internal/repo/postgres"
	"backend/internal/repo/postgres/notice"
	kafkaRepo "backend/internal/repo/kafka"
	"backend/internal/service/dto"
	noticeService "backend/internal/service/notice"
	loggerMock "backend/pkg/logger/mocks"
	"context"
	"fmt"

	// "github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type NoticeSuite struct {
	suite.Suite
}

// func (suite *NoticeSuite) TestSubscriberAddSuccess_01(t provider.T) {
// 	mockLogger := new(loggerMock.Interface)
// 	ctx := context.Background()
// 	cfg := config.Postgres{
// 		Host:     "localhost",
// 		Port:     5438,
// 		User:     "avito",
// 		Password: "avito",
// 		Database: "avito",
// 		Driver:   "postgres",
// 	}
// 	dbConnector, err := postgres.NewDbConn(ctx, &cfg)
// 	if err != nil {
// 		fmt.Println("db is not connected")
// 		return
// 	}
// 	r := postgres.NewRetryAdapter(dbConnector, 2, 2)
// 	noticeR := notice.NewNoticeRepo(mockLogger, dbConnector, r)

// 	noticeS := noticeService.NewNoticeService(mockLogger, noticeR)

// 	mockLogger.On("Errorf", mock.Anything, mock.Anything).Times(0)
// 	mockLogger.On("Infof", mock.Anything, mock.Anything).Times(0)
// 	mockLogger.On("Warnf", mock.Anything, mock.Anything).Times(0)

// 	id, err := uuid.Parse("54930a14-ffab-46e9-ac9f-7f9c50341872")
// 	if err != nil {
// 		fmt.Println("failed to parse")
// 		return
// 	}
// 	houseId, err := uuid.Parse("b5297fe6-a5d9-43a8-9923-b73daabc8ceb")
// 	if err != nil {
// 		fmt.Println("failed to parse")
// 		return
// 	}
// 	response, err := noticeS.Subscribe(ctx, &dto.SubscribeRequest{
// 		UserID:  id,
// 		HouseID: houseId,
// 	})

// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, response)
// }

// func (suite *NoticeSuite) TestCreateNoticesSuccess_01(t provider.T) {
// 	mockLogger := new(loggerMock.Interface)
// 	ctx := context.Background()
	
// 	cfg := config.Postgres{
// 		Host:     "localhost",
// 		Port:     5438,
// 		User:     "avito",
// 		Password: "avito",
// 		Database: "avito",
// 		Driver:   "postgres",
// 	}
// 	dbConnector, err := postgres.NewDbConn(ctx, &cfg)
// 	if err != nil {
// 		fmt.Println("db is not connected")
// 		return
// 	}
// 	r := postgres.NewRetryAdapter(dbConnector, 2, 2)
// 	noticeR := notice.NewNoticeRepo(mockLogger, dbConnector, r)
// 	kafkaR, err := kafkaRepo.NewProducer(mockLogger, []string{"localhost:9092"}, "test")
// 	if err != nil {
// 		fmt.Println("cant connect ot kafka")
// 	}

// 	noticeS := noticeService.NewNoticeService(mockLogger, noticeR, kafkaR)

// 	mockLogger.On("Errorf", mock.Anything, mock.Anything).Times(0)
// 	mockLogger.On("Infof", mock.Anything, mock.Anything).Times(0)
// 	mockLogger.On("Warnf", mock.Anything, mock.Anything).Times(0)

// 	apartmentID, err := uuid.Parse("2569de76-ea61-4b36-bc20-0efb87171681")
// 	if err != nil {
// 		fmt.Println("failed to parse")
// 		return
// 	}
// 	houseID, err := uuid.Parse("b5297fe6-a5d9-43a8-9923-b73daabc8ceb")
// 	if err != nil {
// 		fmt.Println("failed to parse")
// 		return
// 	}
// 	response, err := noticeS.CreateNotices(ctx, &dto.CreateNoticesRequest{
// 		ApartmentID:  apartmentID,
// 		HouseID: houseID,
// 	})

// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, response)
// }

func (suite *NoticeSuite) TestSendNoticesSuccess_01(t provider.T) {
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
	noticeR := notice.NewNoticeRepo(mockLogger, dbConnector, r)
	kafkaR, err := kafkaRepo.NewProducer(mockLogger, []string{"localhost:9093"}, "test")
	if err != nil {
		fmt.Println("cant connect ot kafka")
	}

	noticeS := noticeService.NewNoticeService(mockLogger, noticeR, kafkaR)

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Times(0)
	mockLogger.On("Infof", mock.Anything, mock.Anything).Times(0)
	mockLogger.On("Warnf", mock.Anything, mock.Anything).Times(0)

	// apartmentID, err := uuid.Parse("2569de76-ea61-4b36-bc20-0efb87171681")
	// if err != nil {
	// 	fmt.Println("failed to parse")
	// 	return
	// }
	// houseID, err := uuid.Parse("b5297fe6-a5d9-43a8-9923-b73daabc8ceb")
	// if err != nil {
	// 	fmt.Println("failed to parse")
	// 	return
	// }
	response, err := noticeS.SendNoticesToKafka(ctx, &dto.SendNoticesToKafkaRequest{
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)
}


