package userserver

import (
	"context"
	entityuser "user_service/internal/entity/user"
	usecaseuser "user_service/internal/usecase/user"

	usergen "github.com/DIYORBEKORIFJONOV/chat_service_protos/gen/go/user"
	"google.golang.org/grpc"
)

type userServer struct {
	usergen.UnimplementedUserServiceServer
	user usecaseuser.UserUseCaseIml
}

func RegisterUserServer(GRPCServer *grpc.Server,user usecaseuser.UserUseCaseIml) {
	usergen.RegisterUserServiceServer(GRPCServer,&userServer{
		user: user,
	})
}

func (u *userServer) ChangeUserPassword(ctx context.Context,
	 req *usergen.ChangeUserPasswordRequest) (*usergen.ChangeUserPasswordResponse, error) {
	res,err := u.user.ChangeUserPassword(ctx,&entityuser.ChangeUserPasswordRequest{
		UserID: req.UserId,
		CurrentPassword: req.CurrentPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return &usergen.ChangeUserPasswordResponse{
			Success: res.Success,
		},err
	}
	return &usergen.ChangeUserPasswordResponse{
		Success: res.Success,
	},nil
}

func (u *userServer) CreateUser(ctx  context.Context, req *usergen.CreateUserRequest) (*usergen.UserResponse, error) {
	
	res,err := u.user.CreateUserReq(ctx,&entityuser.CreateUserRequest{
		Name: req.Name,
		Email: req.Email,
		Phone: req.Phone,
		PasswordHash: req.PasswordHash,
		Status: req.Status,
		ProfilePictureURL: req.ProfilePictureUrl,
		Language: req.Language,
	})
	if err != nil {
		return nil,err
	}
	return &usergen.UserResponse{
		User: &usergen.User{
			UserId: res.User.UserID,
			Name: res.User.Name,
			Email: res.User.Email,	
			Phone: res.User.Phone,
			PasswordHash: res.User.PasswordHash,
			Status: res.User.Status,
			ProfilePictureUrl: res.User.ProfilePictureURL,
			Language: res.User.Language,
			CreatedAt: res.User.CreatedAt.String(),
			LastLogin: res.User.LastLogin.String(),
		},
	},nil
}

func (u *userServer) GetAllUser(ctx context.Context, req *usergen.GetAllUserRequest) (*usergen.GetUserAllResponse, error) {
	
	res,err := u.user.GetAllUser(ctx,&entityuser.GetAllUserRequest{
		Page: req.Page,
		Limit: req.Limit,
		Field: req.Field,
		Value: req.Value,
	})
	if err != nil {
		return nil,err
	}
	response := []*usergen.User{}
	for _,user := range res.Users {
		response = append(response, &usergen.User{
			UserId: user.UserID,
			Name: user.Name,
			Email: user.Email,
			Phone: user.Phone,
			PasswordHash: user.PasswordHash,
			Status: user.Status,
			ProfilePictureUrl: user.ProfilePictureURL,
			Language: user.Language,
			CreatedAt: user.CreatedAt.String(),
			LastLogin: user.LastLogin.String(),
		})
	}

	return &usergen.GetUserAllResponse{
		Users: response,
	},nil
}

func (u *userServer)GetUser(ctx context.Context, req *usergen.GetUserRequest) (*usergen.UserResponse, error) {
	
	res,err := u.user.GetUser(ctx,&entityuser.GetUserRequest{
		Field: req.Field,
		Value: req.Value,
	})
	if err != nil{
		return nil,err
	}

	return &usergen.UserResponse{
		User: &usergen.User{
			UserId: res.User.UserID,
			Name: res.User.Name,
			Email: res.User.Email,	
			Phone: res.User.Phone,
			PasswordHash: res.User.PasswordHash,
			Status: res.User.Status,
			ProfilePictureUrl: res.User.ProfilePictureURL,
			Language: res.User.Language,
			CreatedAt: res.User.CreatedAt.String(),
			LastLogin: res.User.LastLogin.String(),
		},
	},nil

}

func (u *userServer) LoginUser(ctx context.Context, req *usergen.LoginUserRequest) (*usergen.LoginUserResponse, error) {
	
	res,err := u.user.LoginUser(ctx,&entityuser.LoginUserRequest{
		Email: req.Email,
		PasswordHash: req.PasswordHash,
	})
	
	if err != nil {
		return &usergen.LoginUserResponse{
			Success: res.Success,
		},err
	}

	return &usergen.LoginUserResponse{
		Success: res.Success,
	},nil
}

func (u *userServer)UpdateUser(ctx context.Context, req *usergen.UpdateUserRequest) (*usergen.UserResponse, error) {
	
	res,err := u.user.UpdateUser(ctx,&entityuser.UpdateUserRequest{
		UserID: req.UserId,
		Name: req.Name,
		Phone: req.Phone,
		Status: req.Status,
		ProfilePictureURL: req.ProfilePictureUrl,
		Language: req.Language,
		LastLogin: req.LastLogin,
	})

	if err != nil {
		return nil,err
	}

	
	return &usergen.UserResponse{
		User: &usergen.User{
			UserId: res.User.UserID,
			Name: res.User.Name,
			Email: res.User.Email,	
			Phone: res.User.Phone,
			PasswordHash: res.User.PasswordHash,
			Status: res.User.Status,
			ProfilePictureUrl: res.User.ProfilePictureURL,
			Language: res.User.Language,
			CreatedAt: res.User.CreatedAt.String(),
			LastLogin: res.User.LastLogin.String(),
		},
	},nil
}