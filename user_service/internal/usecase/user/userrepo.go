package usecaseuser

import (
	"context"
	entityuser "user_service/internal/entity/user"
)

type userRepoUseCase interface {
	SaveUser(ctx context.Context, user *entityuser.User) error
	ChangeUserPassword(ctx context.Context, userId,newPassword string) error
	GetUser(ctx context.Context, field,value string)(*entityuser.User,error)
	GetAllUser(ctx context.Context, req *entityuser.GetAllUserRequest) ([]entityuser.User,error)
	UpdateUser(ctx context.Context, req *entityuser.UpdateUserRequest)(*entityuser.User,error)
}

type UserRepoUseCaseIml struct {
	user userRepoUseCase
}
func NewUserRepoUseCase( user userRepoUseCase) *UserRepoUseCaseIml {
	return &UserRepoUseCaseIml{
		user: user,
	}
}

func (u *UserRepoUseCaseIml)SaveUser(ctx context.Context, user *entityuser.User) error {
	return u.user.SaveUser(ctx,user)
}


func (u *UserRepoUseCaseIml)ChangeUserPassword(ctx context.Context, userId,newPassword string) error {
	return u.user.ChangeUserPassword(ctx,userId,newPassword)
}

func (u *UserRepoUseCaseIml)GetUser(ctx context.Context, field,value string)(*entityuser.User,error) {
	return u.user.GetUser(ctx,field,value)
}

func (u *UserRepoUseCaseIml)GetAllUser(ctx context.Context, req *entityuser.GetAllUserRequest) ([]entityuser.User,error) {
	return u.user.GetAllUser(ctx,req)
}

func (u *UserRepoUseCaseIml)UpdateUser(ctx context.Context, req *entityuser.UpdateUserRequest)(*entityuser.User,error) {
	return u.user.UpdateUser(ctx,req)
}