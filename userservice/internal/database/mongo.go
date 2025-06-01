package database

import (
	accountmodel "account/internal/model"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Database struct {
	M *mongo.Collection
	C context.Context
}

func (db *Database) SaveUser(req *accountmodel.CreateAccountRequest) (*accountmodel.CreateAccountResponse, error) {
	account_id := uuid.NewString()
	hasedPassword, err := Hashing(req.Password)
	if err != nil {
		return nil, err
	}
	account := accountmodel.Account{
		AccountID: account_id,
		Name:      req.Name,
		Phone:     req.Phone,
		Username:  req.Username,
		Password:  hasedPassword,
		Role:      req.Role,
	}
	_, err = db.M.InsertOne(db.C, account)
	if err != nil {
		return nil, err
	}
	return &accountmodel.CreateAccountResponse{
		Account: account,
	}, nil
}
func (db *Database) GetUser(req *accountmodel.GetAccountRequest) (*accountmodel.GetAccountResponse, error) {
	var account accountmodel.Account
	err := db.M.FindOne(db.C, bson.M{"account_id": req.AccountID}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found with this id %v", err)
		}
		return nil, err
	}
	return &accountmodel.GetAccountResponse{
		Account: account,
	}, nil
}
func (db *Database) ChangePassword(req *accountmodel.ChangePasswordRequest) (*accountmodel.ChangePasswordResponse, error) {
	var account accountmodel.Account
	err := db.M.FindOne(db.C, bson.M{"account_id": req.AccountID}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found with this id: %v", req.AccountID)
		}
		return nil, err
	}

	hasedPassword, err := Hashing(req.NewPassword)
	if err != nil {
		return nil, err
	}

	_, err = db.M.UpdateOne(
		db.C,
		bson.M{"account_id": req.AccountID},
		bson.M{"$set": bson.M{"password": hasedPassword}},
	)
	if err != nil {
		return nil, err
	}

	return &accountmodel.ChangePasswordResponse{Changed: true}, nil
}

func (db *Database) DeleteUser(req *accountmodel.DeleteUserRequest) (*accountmodel.DeleteUserResponse, error) {
	_, err := db.M.DeleteOne(db.C, bson.M{"account_id": req.AccountID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found with this id %v", err)
		}
		return nil, err
	}
	return &accountmodel.DeleteUserResponse{Deleted: true}, nil
}
func (db *Database) UserExists(req *accountmodel.UserExistsRequest) (*accountmodel.UserExistsResponse, error) {
	var account accountmodel.Account
	err := db.M.FindOne(db.C, bson.M{"phone": req.Phone}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &accountmodel.UserExistsResponse{Exists: false}, nil
		}
		return nil, err
	}
	if account.Phone != req.Phone {
		return &accountmodel.UserExistsResponse{Exists: false}, nil
	}
	return &accountmodel.UserExistsResponse{Exists: true}, nil

}

func (db *Database) Login(req *accountmodel.LoginRequest) (*accountmodel.GetAccountResponse, error) {
	var account accountmodel.Account
	err := db.M.FindOne(db.C, bson.M{"username": req.Username}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found with this username %v", err)
		}
		return nil, err
	}
	err = CompareHashAndPassword(account.Password, req.Password)
	if err != nil {
		return nil, fmt.Errorf("incorrect password %v", err)
	}
	return &accountmodel.GetAccountResponse{
		Account: account,
	}, nil
}

func Hashing(password string) (string, error) {
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptedPassword), nil
}
func CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
