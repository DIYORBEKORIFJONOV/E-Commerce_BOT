package service

import (
	__Account "account/internal/protos/.Account"
	"account/internal/service/adjust"
	"context"
)

type Service struct {
	__Account.UnimplementedAccountServiceServer
	A *adjust.AdjustService
}

func (s *Service) ChangePassword(ctx context.Context, req *__Account.ChangePasswordRequest) (*__Account.ChangePasswordResponse, error) {
	res, err := s.A.ChangePassword(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) CreateAccount(ctx context.Context, req *__Account.CreateAccountRequest) (*__Account.CreateAccountResponse, error) {
	res, err := s.A.SaveUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) DeleteUser(ctx context.Context, req *__Account.DeleteUserRequest) (*__Account.DeleteUserResponse, error) {
	res, err := s.A.DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) GetAccount(ctx context.Context, req *__Account.GetAccountRequest) (*__Account.GetAccountResponse, error) {
	res, err := s.A.GetUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) UserExists(ctx context.Context, req *__Account.UserExistsRequest) (*__Account.UserExistsResponse, error) {
	res, err := s.A.UserExists(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Login(ctx context.Context, req *__Account.LoginRequest) (*__Account.GetAccountResponse, error) {
	res, err := s.A.Login(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
