package usecaseuser

import (
	"context"
	entityuser "user_service/internal/entity/user"
)

type userUseCase interface { 
	CreateUserReq(ctx context.Context, req *entityuser.CreateUserRequest) (*entityuser.UserResponse,error)
	ChangeUserPassword(ctx context.Context, req *entityuser.ChangeUserPasswordRequest)(*entityuser.ChangeUserPasswordResponse,error)
	GetUser(ctx context.Context, req *entityuser.GetUserRequest)(*entityuser.UserResponse,error)
	GetAllUser(ctx context.Context, req *entityuser.GetAllUserRequest) (*entityuser.GetUserAllResponse,error)
	LoginUser(ctx context.Context, req *entityuser.LoginUserRequest) (*entityuser.LoginUserResponse,error)
	UpdateUser(ctx context.Context, req *entityuser.UpdateUserRequest) (*entityuser.UserResponse,error)
}

type UserUseCaseIml struct {
	user userUseCase
}

func NewUserUseCase(user userUseCase) *UserUseCaseIml {
	return &UserUseCaseIml{
		user:user,
	}
}

func (u *UserUseCaseIml)CreateUserReq(ctx context.Context, req *entityuser.CreateUserRequest) (*entityuser.UserResponse,error) {
	return u.user.CreateUserReq(ctx,req)
}

func (u *UserUseCaseIml)ChangeUserPassword(ctx context.Context, req *entityuser.ChangeUserPasswordRequest)(*entityuser.ChangeUserPasswordResponse,error) {
	return u.user.ChangeUserPassword(ctx,req)
}
func (u *UserUseCaseIml)GetUser(ctx context.Context, req *entityuser.GetUserRequest)(*entityuser.UserResponse,error) {
	return u.user.GetUser(ctx,req)
}

func (u *UserUseCaseIml)GetAllUser(ctx context.Context, req *entityuser.GetAllUserRequest) (*entityuser.GetUserAllResponse,error) {
	return u.user.GetAllUser(ctx,req)
}

func (u *UserUseCaseIml)LoginUser(ctx context.Context, req *entityuser.LoginUserRequest) (*entityuser.LoginUserResponse,error) {
	return u.user.LoginUser(ctx,req)
}

func (u *UserUseCaseIml)UpdateUser(ctx context.Context, req *entityuser.UpdateUserRequest) (*entityuser.UserResponse,error) {
	return u.user.UpdateUser(ctx,req)
}