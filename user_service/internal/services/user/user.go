package userservice

import (
	"context"
	"errors"
	"time"
	entityuser "user_service/internal/entity/user"
	userRedis "user_service/internal/infastructure/repository/redis/user"
	usecaseuser "user_service/internal/usecase/user"
	logger "user_service/log"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type logContext struct {
	console   *zerolog.Logger
	logFile   *zerolog.Logger
	operation string
	userID    string
	email     string
}

func (l *logContext) info(msg string) {
	logEvent := l.console.Info()
	if l.userID != "" {
		logEvent.Str("userId", l.userID)
	}
	if l.email != "" {
		logEvent.Str("email", l.email)
	}
	logEvent.Msg(msg)

	logEvent = l.logFile.Info()
	if l.userID != "" {
		logEvent.Str("userId", l.userID)
	}
	if l.email != "" {
		logEvent.Str("email", l.email)
	}
	logEvent.Msg(msg)
}

func (l *logContext) error(err error, msg string) {
	l.console.Err(err).Str("operation", l.operation).Msg(msg)
	l.logFile.Err(err).Str("operation", l.operation).Msg(msg)
}

func (l *logContext) warn(msg string) {
	l.console.Warn().Str("operation", l.operation).Msg(msg)
	l.logFile.Warn().Str("operation", l.operation).Msg(msg)
}

type UserService struct {
	user   *usecaseuser.UserRepoUseCaseIml
	logger *logger.Loggger
	kash *userRedis.RedisUserRepository
}

func NewUserService(user *usecaseuser.UserRepoUseCaseIml, logger *logger.Loggger,kash *userRedis.RedisUserRepository) *UserService {
	return &UserService{
		user:   user,
		logger: logger,
		kash: kash,
	}
}

func (u *UserService) newLogContext(operation string) *logContext {
	console := u.logger.Console.With().Str("operation", operation).Logger()
	logFile := u.logger.Logger.With().Str("operation", operation).Logger()

	return &logContext{
		console:   &console,
		logFile:   &logFile,
		operation: operation,
	}
}

func (u *UserService) CreateUserReq(ctx context.Context, req *entityuser.CreateUserRequest) (*entityuser.UserResponse, error) {
	logs := u.newLogContext("service.CreateUserReq")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		logs.error(err, "failed to hash password")
		return nil, err
	}

	user := &entityuser.User{
		UserID:            uuid.NewString(),
		Name:             req.Name,
		Email:            req.Email,
		Phone:            req.Phone,
		PasswordHash:     string(passwordHash),
		Status:           req.Status,
		ProfilePictureURL: req.ProfilePictureURL,
		Language:         req.Language,
		CreatedAt:        time.Now(),
		LastLogin:        time.Now(),
	}

	logs.userID = user.UserID
	logs.email = user.Email
	logs.info("starting to save the user to the database")

	if err := u.user.SaveUser(ctx, user); err != nil {
		logs.error(err, "failed saving user")
		return nil, err
	}

	logs.info("user successfully saved to the database")
	return &entityuser.UserResponse{User: *user}, nil
}

func (u *UserService) ChangeUserPassword(ctx context.Context, req *entityuser.ChangeUserPasswordRequest) (*entityuser.ChangeUserPasswordResponse, error) {
	logs := u.newLogContext("service.ChangeUserPassword")
	logs.userID = req.UserID

	logs.info("starting password change process")

	user, err := u.GetUser(ctx, &entityuser.GetUserRequest{
		Field: "user_id",
		Value: req.UserID,
	})
	if err != nil {
		logs.error(err, "failed to fetch user for password change")
		return &entityuser.ChangeUserPasswordResponse{Success: false}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.User.PasswordHash), []byte(req.CurrentPassword)); err != nil {
		logs.warn("invalid current password provided")
		return &entityuser.ChangeUserPasswordResponse{Success: false}, errors.New("invalid current password")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logs.error(err, "failed to hash new password")
		return nil, err
	}

	err = u.user.ChangeUserPassword(ctx, req.UserID, string(passwordHash))
	if err != nil {
		logs.error(err, "failed to update user password in database")
		return &entityuser.ChangeUserPasswordResponse{Success: false}, err
	}

	logs.info("password successfully updated")

	err = u.kash.DeleteFromKash(ctx,req.UserID)
	if err != nil {
		logs.error(err,"deleteing from kash")
		return nil,err
	}
	return &entityuser.ChangeUserPasswordResponse{Success: true}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *entityuser.GetUserRequest) (*entityuser.UserResponse, error) {
	logs := u.newLogContext("service.GetUser")

	if req.Field == "user_id" {
		logs.info("Cheking in kash")
		kashUser,err := u.kash.GetFromKash(ctx,req.Value)
		if err != nil {
			logs.error(err,"kash user")
			return nil,err
		}
		if kashUser != nil {
			return &entityuser.UserResponse{
				User: *kashUser,
			},nil
		}
	}
	
	logs.info("starting user fetch process")
	user, err := u.user.GetUser(ctx, req.Field, req.Value)
	if err != nil {
		logs.error(err, "failed to fetch user from database")
		return nil, err
	}

	logs.userID = user.UserID
	logs.info("user fetched successfully")

	logs.info("adding to kash")
	err = u.kash.SaveToKash(ctx,user)
	if err != nil {
		logs.warn("adding to kash")
	}
	return &entityuser.UserResponse{User: *user}, nil
}

func (u *UserService) GetAllUser(ctx context.Context, req *entityuser.GetAllUserRequest) (*entityuser.GetUserAllResponse, error) {
	logs := u.newLogContext("service.GetAllUser")

	logs.info("starting fetch process for all users")

	users, err := u.user.GetAllUser(ctx, req)
	if err != nil {
		logs.error(err, "failed to fetch all users from database")
		return nil, err
	}

	logs.info("fetched all users successfully")
	return &entityuser.GetUserAllResponse{Users: users}, nil
}

func (u *UserService) LoginUser(ctx context.Context, req *entityuser.LoginUserRequest) (*entityuser.LoginUserResponse, error) {
	logs := u.newLogContext("service.LoginUser")
	logs.email = req.Email

	logs.info("starting login process")

	user, err := u.user.GetUser(ctx, "email", req.Email)
	if err != nil {
		logs.error(err, "user not found during login")
		return &entityuser.LoginUserResponse{Success: false}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.PasswordHash)); err != nil {
		logs.warn("invalid password during login")
		return &entityuser.LoginUserResponse{Success: false}, errors.New("invalid password")
	}

	logs.userID = user.UserID
	logs.info("user logged in successfully")
	return &entityuser.LoginUserResponse{Success: true}, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *entityuser.UpdateUserRequest) (*entityuser.UserResponse, error) {
	logs := u.newLogContext("service.UpdateUser")
	logs.userID = req.UserID

	logs.info("starting user update process")

	user, err := u.user.UpdateUser(ctx, req)
	if err != nil {
		logs.error(err, "failed to update user")
		return nil, err
	}

	logs.info("user updated successfully")

	err = u.kash.DeleteFromKash(ctx,req.UserID)
	if err != nil {
		logs.error(err,"deleteing from kash")
		return nil,err
	}
	return &entityuser.UserResponse{User: *user}, nil
}