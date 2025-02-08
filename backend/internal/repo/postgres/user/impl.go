package user

import (
	repoDto "backend/internal/repo/dto"
	"backend/internal/repo/postgres"
	"backend/pkg/logger"
	"context"
)

type UserRepo struct {
	logger       logger.Interface
	retryAdapter postgres.RetryAdapterIntf
}

func NewUserRepo(logger logger.Interface, retryAdapter postgres.RetryAdapterIntf) *UserRepo {
	return &UserRepo{
		logger:       logger,
		retryAdapter: retryAdapter,
	}
}

func (r *UserRepo) Add(ctx context.Context, request *repoDto.AddUserRequest) (*repoDto.AddUserResponse, error) {
	method := "UserRepo -- Add"
	r.logger.Infof("%s", method)

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "insert into users(creation_time, email, password, is_moderator, totp_secret) values ($1, $2, $3, $4, $5) returning id"

	response := repoDto.AddUserResponse{}
	rows := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.CreationTime,
		request.Email,
		request.Password,
		request.IsModerator,
		request.TotpSecret,
	)
	defer rows.Close()
	err := rows.Scan(
		response.ID,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrExec, err)
		return nil, ErrQueringRow
	}

	return &response, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, request *repoDto.GetUserByEmailRequest) (*repoDto.GetUserByEmailResponse, error) {
	method := "UserRepo -- GetByEmail"
	r.logger.Infof("%s", method)

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "select id, creation_time, email, password, is_moderator, totp_secret from users where email = $1"

	response := repoDto.GetUserByEmailResponse{}
	err := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.Email,
	).Scan(
		response.ID,
		response.CreationTime,
		response.Email,
		response.Password,
		response.TotpSecret,
		response.IsModerator,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrQueringRow)
		return nil, ErrQueringRow
	}

	return &response, nil
}
