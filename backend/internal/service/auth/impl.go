package auth

import (
	"backend/internal/model"
	repoDto "backend/internal/repo/dto"
	userRepo "backend/internal/repo/user"
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
		s.logger.Errorf("%s -- %s", method, model.ErrUser_NilContext)
		return nil, model.ErrUser_NilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrUser_NilRequest)
		return nil, model.ErrUser_NilRequest
	}

	if !validator.IsValidEmail(request.Email) {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_BadMail)
		return nil, model.ErrUser_BadMail
	}

	if request.Password == "" {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_BadPassword)
		return nil, model.ErrUser_BadPassword
	}

	if !(request.Role == "client" || request.Role == "moderator") {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_BadType)
		return nil, model.ErrUser_BadType
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_ErrorHashPassword)
		return nil, model.ErrUser_ErrorHashPassword
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
		s.logger.Warnf("%s -- %s", method, model.ErrUser_ErrorTotpGenerate)
		return nil, model.ErrUser_ErrorTotpGenerate
	}

	totpSecretEncypt, err := aes.AesEncrypt(aes.KEY, []byte(totp.Secret()))
	if err != nil {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_TotpEncrypt)
		return nil, model.ErrUser_TotpEncrypt
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
		CreationTime: response.CreationTime,
		TotpSecret:   totp.Secret(),
	}, nil
}

func (s *AuthService) LogIn(ctx context.Context, request *serviceDto.LogInRequest) (*serviceDto.LogInResponse, error) {
	method := "AuthService -- LogIn"
	s.logger.Infof("%s", method)

	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrUser_NilContext)
		return nil, model.ErrUser_NilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrUser_NilRequest)
		return nil, model.ErrUser_NilRequest
	}

	if !validator.IsValidEmail(request.Email) {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_BadMail)
		return nil, model.ErrUser_BadMail
	}

	if request.Password == "" {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_BadPassword)
		return nil, model.ErrUser_BadPassword
	}

	user, err := s.UserRepoIntf.GetByEmail(ctx, &repoDto.GetUserByEmailRequest{
		Email: request.Email,
	})

	if err != nil {
		return nil, err
	}

	if request.Email != user.Email {
		s.logger.Errorf("%s -- %s", method, model.ErrUser_BadMail)
		return nil, model.ErrUser_BadMail
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password))
	if err != nil {
		s.logger.Errorf("%s -- %s", method, model.ErrUser_TotpDecrypt)
		return nil, model.ErrUser_TotpDecrypt
	}

	totpSecret, err := aes.AesDecrypt(aes.KEY, user.TotpSecret)
	if err != nil {
		s.logger.Errorf("%s -- %s", method, model.ErrUser_TotpDecrypt)
		return nil, model.ErrUser_TotpDecrypt
	}

	if !totp.Validate(request.Token, string(totpSecret)) {
		s.logger.Warnf("%s -- %s", method, model.ErrUser_BadToken)
		return nil, model.ErrUser_BadToken
	}

	// pasetoStruct, err := paseto.NewPaseto(paseto.KEY)
	// if err != nil {
	// 	a.logger.Errorf("failed to generate paseto struct", err)
	// 	return "", err
	// }
	token, err := s.pasetoToken.CreateToken(user.ID, user.IsModerator, 1*time.Hour)
	if err != nil {
		s.logger.Warnf("%s -- %s", method, err)
		return nil, err
	}

	return &serviceDto.LogInResponse{
		Token: token,
	}, nil
}
