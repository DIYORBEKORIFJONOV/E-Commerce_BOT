package interfaceaccount

import accountmodel "account/internal/model"

type InterfaceAccount interface {
	SaveUser(req *accountmodel.CreateAccountRequest) (*accountmodel.CreateAccountResponse, error)
	GetUser(req *accountmodel.GetAccountRequest) (*accountmodel.GetAccountResponse, error)
	ChangePassword(req *accountmodel.ChangePasswordRequest) (*accountmodel.ChangePasswordResponse, error)
	DeleteUser(req *accountmodel.DeleteUserRequest) (*accountmodel.DeleteUserResponse, error)
	UserExists(req *accountmodel.UserExistsRequest) (*accountmodel.UserExistsResponse, error)
	Login(req *accountmodel.LoginRequest) (*accountmodel.GetAccountResponse, error)
}
