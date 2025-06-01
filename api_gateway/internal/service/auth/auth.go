package auth

import (
	"context"
	"fmt"
	authentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/auth"
	clientgrpcserver "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/client_grpc_server"
	redisCash "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/infastructure/redis"
	tokensettings "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/token"
	accountgen "github.com/diyorbek/E-Commerce_BOT/api_gateway/pkg/protos/gen/account"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/until"
)

type AuthService struct {
	redisCash    redisCash.Service
	phoneSetting *until.PhoneSetting
	grpcClient   clientgrpcserver.ServiceClient
}

func NewAuthService(redisCash redisCash.Service, phoneSetting *until.PhoneSetting, grpcClient clientgrpcserver.ServiceClient) *AuthService {
	return &AuthService{
		redisCash:    redisCash,
		phoneSetting: phoneSetting,
		grpcClient:   grpcClient,
	}
}

func (s *AuthService) RegisterAccount(ctx context.Context, account *authentity.RegisterForm) (response *authentity.RegisterResponse, err error) {
	response = new(authentity.RegisterResponse)

	switch existResponse, err := s.grpcClient.AccountService().UserExists(ctx, &accountgen.UserExistsRequest{
		Phone: account.Phone,
	}); {
	case err != nil:
		return nil, err
	case existResponse.Exists:
		response.AlreadyExists = true
		return nil, err
	default:
		response.AlreadyExists = false
	}

	//TODO use  secretCode
	//secretCode := s.phoneSetting.Generate4DigitCode()

	err = s.phoneSetting.SendSMS(account.Phone[1:], "Bu Eskiz dan test")
	if err != nil {
		response.Sent = false
		return
	}

	err = s.redisCash.SaveAccount(ctx, &redisCash.SaveAccount{
		Name:       account.Name,
		Phone:      account.Phone,
		Password:   account.Password,
		SenderCode: "Bu Eskiz dan test",
		Username:   account.Username,
	})
	if err != nil {
		response.Sent = false
		return
	}

	response.Sent = true
	return
}

func (s *AuthService) VerifyAccount(ctx context.Context, request *authentity.VerifyRequest) (token string, err error) {
	cashAccount, err := s.redisCash.GetAccount(ctx, request.Phone)
	if err != nil {
		return
	}
	if cashAccount.SenderCode != request.SecretCode {
		err = fmt.Errorf("sender code not match")
		return
	}

	account, err := s.grpcClient.AccountService().CreateAccount(ctx, &accountgen.CreateAccountRequest{
		Phone:    cashAccount.Phone,
		Name:     cashAccount.Name,
		Username: cashAccount.Username,
		Password: cashAccount.Password,
	})
	if err != nil {
		return
	}
	return tokensettings.GenerateToken(authentity.Account{
		Id:           account.Account.AccountId,
		Name:         account.Account.Name,
		Username:     account.Account.Username,
		Phone:        account.Account.Phone,
		PasswordHash: account.Account.Password,
		Role:         account.Account.Role,
	})
}

func (s *AuthService) Login(ctx context.Context, username string, password string) (token string, err error) {
	account, err := s.grpcClient.AccountService().Login(ctx, &accountgen.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return
	}
	return tokensettings.GenerateToken(authentity.Account{
		Id:           account.Account.AccountId,
		Name:         account.Account.Name,
		Username:     account.Account.Username,
		Phone:        account.Account.Phone,
		PasswordHash: account.Account.Password,
		Role:         account.Account.Role,
	})
}

func (s *AuthService) ChangePassword(ctx context.Context, userId string, password string) (changed bool, err error) {
	response, err := s.grpcClient.AccountService().ChangePassword(ctx, &accountgen.ChangePasswordRequest{
		AccountId:   userId,
		NewPassword: password,
	})

	if err != nil {
		return
	}
	changed = response.Changed
	return
}
