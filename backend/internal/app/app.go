package app

import (
	"backend/config"
	"backend/internal/repo/postgres"
	apartmentRepo "backend/internal/repo/postgres/apartment"
	houseRepo "backend/internal/repo/postgres/house"
	noticetRepo "backend/internal/repo/postgres/notice"
	userRepo "backend/internal/repo/postgres/user"

	apartmentService "backend/internal/service/apartment"
	authService "backend/internal/service/auth"
	houseService "backend/internal/service/house"
	noticeService "backend/internal/service/notice"

	"backend/pkg/logger"
	"backend/pkg/token/paseto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	authService      authService.AuthIntf
	houseService     houseService.HouseIntf
	apartmentService apartmentService.ApartmentIntf
	noticeService    noticeService.NoticeIntf
}

func NewApp(
	logger logger.Interface,
	dbConnector *pgxpool.Pool,
	pasetoStruct *paseto.Paseto,
	cfg *config.Config,
) *App {
	retryAdapter := postgres.NewRetryAdapter(dbConnector, cfg.DB.ExecRetryNumber, cfg.DB.SleepTimeMilsec)

	userR := userRepo.NewUserRepo(logger, retryAdapter)
	houseR := houseRepo.NewHouseRepo(logger, retryAdapter)
	apartmentR := apartmentRepo.NewApartmentRepo(logger, retryAdapter)
	noticeR := noticetRepo.NewNoticeRepo(logger, dbConnector, retryAdapter)

	authS := authService.NewAuthService(pasetoStruct, logger, userR)
	houseS := houseService.NewHouseService(logger, houseR)
	apartmentS := apartmentService.NewApartmentService(logger, apartmentR)
	noticeS := noticeService.NewNoticeService(logger, noticeR)

	return &App{
		authService:      authS,
		houseService:     houseS,
		apartmentService: apartmentS,
		noticeService:    noticeS,
	}
}
