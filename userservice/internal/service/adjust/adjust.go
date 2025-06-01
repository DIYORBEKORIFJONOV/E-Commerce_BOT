package adjust

import (
	interfaceaccount "account/internal/interface"
	accountmodel "account/internal/model"
	__Account "account/internal/protos/.Account"
)

type AdjustService struct {
	I interfaceaccount.InterfaceAccount
}

func (u *AdjustService) SaveUser(req *__Account.CreateAccountRequest) (*__Account.CreateAccountResponse, error) {
	var account = accountmodel.CreateAccountRequest{
		Name:     req.Name,
		Phone:    req.Phone,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}
	resp, err := u.I.SaveUser(&account)
	if err != nil {
		return nil, err
	}
	return &__Account.CreateAccountResponse{
		Account: &__Account.Account{
			AccountId: resp.Account.AccountID,
			Name:      resp.Account.Name,
			Phone:     resp.Account.Phone,
			Username:  resp.Account.Username,
			Password:  resp.Account.Password,
			Role:      resp.Account.Role,
		},
	}, nil
}
func (u *AdjustService) GetUser(req *__Account.GetAccountRequest) (*__Account.GetAccountResponse, error) {
	var account = accountmodel.GetAccountRequest{
		AccountID: req.AccountId,
	}
	resp, err := u.I.GetUser(&account)
	if err != nil {
		return nil, err
	}
	return &__Account.GetAccountResponse{
		Account: &__Account.Account{
			AccountId: resp.Account.AccountID,
			Name:      resp.Account.Name,
			Phone:     resp.Account.Phone,
			Username:  resp.Account.Username,
			Password:  resp.Account.Password,
			Role:      resp.Account.Role,
		},
	}, nil
}
func (u *AdjustService) ChangePassword(req *__Account.ChangePasswordRequest) (*__Account.ChangePasswordResponse, error) {
	var account = accountmodel.ChangePasswordRequest{
		AccountID:   req.AccountId,
		NewPassword: req.NewPassword,
	}
	resp, err := u.I.ChangePassword(&account)
	if err != nil {
		return nil, err
	}
	return &__Account.ChangePasswordResponse{
		Changed: resp.Changed,
	}, nil
}

func (u *AdjustService) DeleteUser(req *__Account.DeleteUserRequest) (*__Account.DeleteUserResponse, error) {
	var account = accountmodel.DeleteUserRequest{
		AccountID: req.AccountId,
	}
	resp, err := u.I.DeleteUser(&account)
	if err != nil {
		return nil, err
	}
	return &__Account.DeleteUserResponse{
		Deleted: resp.Deleted,
	}, nil
}
func (u *AdjustService) UserExists(req *__Account.UserExistsRequest) (*__Account.UserExistsResponse, error) {
	var account = accountmodel.UserExistsRequest{
		Phone: req.Phone,
	}
	resp, err := u.I.UserExists(&account)
	if err != nil {
		return nil, err
	}
	return &__Account.UserExistsResponse{
		Exists: resp.Exists,
	}, nil
}

func (u *AdjustService) Login(req *__Account.LoginRequest) (*__Account.GetAccountResponse, error) {
	var account = accountmodel.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	resp, err := u.I.Login(&account)
	if err != nil {
		return nil, err
	}
	return &__Account.GetAccountResponse{
		Account: &__Account.Account{
			AccountId: resp.Account.AccountID,
			Name:      resp.Account.Name,
			Phone:     resp.Account.Phone,
			Username:  resp.Account.Username,
			Password:  resp.Account.Password,
			Role:      resp.Account.Role,
		},
	}, nil
}
