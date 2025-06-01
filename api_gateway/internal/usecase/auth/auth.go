package authusecase

import (
	"context"
	authentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/auth"
)

type authUseCaseInterface interface {
	RegisterAccount(ctx context.Context, request *authentity.RegisterForm) (response *authentity.RegisterResponse, err error)
	VerifyAccount(ctx context.Context, request *authentity.VerifyRequest) (token string, err error)
	Login(ctx context.Context, username string, password string) (token string, err error)
	ChangePassword(ctx context.Context, userId string, password string) (changed bool, err error)
}

type AuthUseCaseIml struct {
	authUseCaseInterface
}

func NewAuthUseCaseIml(authUseCaseInterface authUseCaseInterface) *AuthUseCaseIml {
	return &AuthUseCaseIml{
		authUseCaseInterface: authUseCaseInterface,
	}
}

func (a *AuthUseCaseIml) ChangePassword(ctx context.Context, userId string, password string) (changed bool, err error) {
	return a.ChangePassword(ctx, userId, password)
}

func (a *AuthUseCaseIml) RegisterAccount(ctx context.Context, request *authentity.RegisterForm) (response *authentity.RegisterResponse, err error) {
	return a.authUseCaseInterface.RegisterAccount(ctx, request)
}

func (a *AuthUseCaseIml) VerifyAccount(ctx context.Context, request *authentity.VerifyRequest) (token string, err error) {
	return a.authUseCaseInterface.VerifyAccount(ctx, request)
}

func (a *AuthUseCaseIml) Login(ctx context.Context, username string, password string) (token string, err error) {
	return a.authUseCaseInterface.Login(ctx, username, password)
}
