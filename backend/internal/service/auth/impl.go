package auth

import (
	repoDto "backend/internal/repo/dto"
	userRepo "backend/internal/repo/postgres/user"
	serviceDto "backend/internal/service/dto"
	"backend/pkg/aes"
	"backend/pkg/logger"
	"backend/pkg/token/paseto"
	"context"
	"time"
	"backend/pkg/validator"

	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	pasetoToken  paseto.Interface
	logger       logger.Interface
	UserRepoIntf userRepo.UserIntf
}

func NewAuthService(pasetoToken paseto.Interface, logger logger.Interface, UserRepoIntf userRepo.UserIntf) *AuthService {
	return &AuthService{
		pasetoToken:  pasetoToken,
		logger:       logger,
		UserRepoIntf: UserRepoIntf,
	}
}

func (s *AuthService) SignUp(ctx context.Context, request *serviceDto.SignUpRequest) (*serviceDto.SignUpResponse, error) {
	method := "AuthServie -- SignUp"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	if !validator.IsValidEmail(request.Email) {
		s.logger.Warnf("%s -- %s", method, ErrBadMail)
		return nil, ErrBadMail
	}

	if request.Password == "" {
		s.logger.Warnf("%s -- %s", method, ErrBadPassword)
		return nil, ErrBadPassword
	}

	if !(request.Role == "client" || request.Role == "moderator") {
		s.logger.Warnf("%s -- %s", method, ErrBadType)
		return nil, ErrBadType
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Warnf("%s -- %s", method, ErrErrorHashPassword)
		return nil, ErrErrorHashPassword
	}

	var isModerator bool
	if request.Role == "moderator" {
		isModerator = true
	} else {
		isModerator = false
	}

	totp, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "avito-bootcamp",
		AccountName: request.Email,
	})
	if err != nil {
		s.logger.Warnf("%s -- %s", method, ErrErrorTotpGenerate)
		return nil, ErrErrorTotpGenerate
	}

	totpSecretEncypt, err := aes.AesEncrypt(aes.KEY, []byte(totp.Secret()))
	if err != nil {
		s.logger.Warnf("%s -- %s", method, ErrTotpEncrypt)
		return nil, ErrTotpEncrypt
	}

	response, err := s.UserRepoIntf.Add(ctx, &repoDto.AddUserRequest{
		CreationTime: time.Now(),
		Email:        request.Email,
		Password:     hashPassword,
		IsModerator:  isModerator,
		TotpSecret:   totpSecretEncypt,
	})

	if err != nil {
		return nil, err
	}
	return &serviceDto.SignUpResponse{
		ID:           response.ID,
		TotpSecret:   totp.Secret(),
	}, nil
}

func (s *AuthService) LogIn(ctx context.Context, request *serviceDto.LogInRequest) (*serviceDto.LogInResponse, error) {
	method := "AuthService -- LogIn"
	s.logger.Infof("%s", method)

	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	if !validator.IsValidEmail(request.Email) {
		s.logger.Warnf("%s -- %s", method, ErrBadMail)
		return nil, ErrBadMail
	}

	if request.Password == "" {
		s.logger.Warnf("%s -- %s", method, ErrBadPassword)
		return nil, ErrBadPassword
	}

	user, err := s.UserRepoIntf.GetByEmail(ctx, &repoDto.GetUserByEmailRequest{
		Email: request.Email,
	})

	if err != nil {
		return nil, err
	}

	if request.Email != user.Email {
		s.logger.Errorf("%s -- %s", method, ErrBadMail)
		return nil, ErrBadMail
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password))
	if err != nil {
		s.logger.Errorf("%s -- %s", method, ErrTotpDecrypt)
		return nil, ErrTotpDecrypt
	}

	totpSecret, err := aes.AesDecrypt(aes.KEY, user.TotpSecret)
	if err != nil {
		s.logger.Errorf("%s -- %s", method, ErrTotpDecrypt)
		return nil, ErrTotpDecrypt
	}

	if !totp.Validate(request.Token, string(totpSecret)) {
		s.logger.Warnf("%s -- %s", method, ErrBadToken)
		return nil, ErrBadToken
	}

	token, err := s.pasetoToken.CreateToken(user.ID, user.IsModerator, 1*time.Hour)
	if err != nil {
		s.logger.Warnf("%s -- %s", method, err)
		return nil, err
	}

	return &serviceDto.LogInResponse{
		Token: token,
	}, nil
}
