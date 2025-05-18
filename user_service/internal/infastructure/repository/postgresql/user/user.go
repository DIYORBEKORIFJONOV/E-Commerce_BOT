package postgresuser

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	entityuser "user_service/internal/entity/user"
	"user_service/internal/infastructure/postgres"
)


type UserRepository struct {
	db *postgres.PostgresDB
	tableName string
}


func NewUserRepository(db *postgres.PostgresDB) *UserRepository {
	return &UserRepository{
		db: db,
		tableName: "users",
	}
}



//select query
func selectQuery() string {
    return `user_id, name, email, phone, password_hash, status, profile_picture_url, language, created_at, last_login`
}


func (u *UserRepository)SaveUser(ctx context.Context, user *entityuser.User) error {

	data := map[string]interface{} {
		"user_id":user.UserID,
		"name":user.Name,
		"email":user.Email,
		"phone":user.Phone,
		"password_hash":user.PasswordHash,
		"status":user.Status,
		"profile_picture_url":user.ProfilePictureURL,
		"language":user.Language,
		"created_at":user.CreatedAt,
		"last_login":user.LastLogin,
	}

	query,args,err := u.db.Sq.Builder.Insert(u.tableName).SetMap(data).ToSql()
	if err != nil {
		return err
	}

	_,err = u.db.Exec(ctx,query,args...)
	if err != nil {
		return err
	}

	return nil
}


func (u *UserRepository)ChangeUserPassword(ctx context.Context, userId,newPassword string) error {

	data := map[string]interface{}{
		"password_hash":newPassword,
	}

	query,args,err := u.db.Sq.Builder.Update(u.tableName).SetMap(data).
	Where(u.db.Sq.Equal("user_id",userId)).ToSql()
	if err != nil {
		return err
	}

	_,err = u.db.Exec(ctx,query,args...)
	if err != nil {
		return err
	}

	return nil
}


func (u *UserRepository)GetUser(ctx context.Context, field,value string)(*entityuser.User,error) {

	query,args,err := u.db.Sq.Builder.Select(selectQuery()).
	From(u.tableName).Where(u.db.Sq.Equal(field,value)).ToSql()
	if err!= nil {
		return nil,err
	}
	var userResponse entityuser.User
	var lastLogin sql.NullTime

	err = u.db.QueryRow(ctx,query,args...).Scan(
		&userResponse.UserID,
		&userResponse.Name,
		&userResponse.Email,
		&userResponse.Phone,
		&userResponse.PasswordHash,
		&userResponse.Status,
		&userResponse.ProfilePictureURL,
		&userResponse.Language,
		&userResponse.CreatedAt,
		&lastLogin,
	)
	if err != nil {
		return nil,err
	}
	if !lastLogin.Time.IsZero() {
		userResponse.LastLogin = lastLogin.Time
	}
	return &userResponse,nil
}

func (u *UserRepository)GetAllUser(ctx context.Context, req *entityuser.GetAllUserRequest) ([]entityuser.User,error) {

	toSql := u.db.Sq.Builder.Select(selectQuery()).From(u.tableName)
	if req.Field != "" && req.Value != "" {
		toSql = toSql.Where(u.db.Sq.Equal(req.Field,req.Value))
	}
	if req.Limit != 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}
	if req.Page != 0 {
		toSql = toSql.Offset(uint64(req.Page))
	}
	
	query,args,err := toSql.ToSql()
	if err != nil {
		return []entityuser.User{},nil
	}

	rows,err := u.db.Query(ctx,query,args...)
	if err != nil {
		return []entityuser.User{},nil
	}

	defer rows.Close()

	var response []entityuser.User

	for rows.Next() {
		var userResponse entityuser.User
		err = rows.Scan(
			&userResponse.UserID,
			&userResponse.Name,
			&userResponse.Email,
			&userResponse.Phone,
			&userResponse.PasswordHash,
			&userResponse.Status,
			&userResponse.ProfilePictureURL,
			&userResponse.Language,
			&userResponse.CreatedAt,
			&userResponse.LastLogin,
		)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		response = append(response, userResponse)
	}

	return response,nil
}


func (u *UserRepository)UpdateUser(ctx context.Context, req *entityuser.UpdateUserRequest)(*entityuser.User,error) {
	data := map[string]interface{}{}
	if req.Language != "" {
		data["language"] = req.Language
	}
	if req.LastLogin != "" {
		data["last_login"] = req.LastLogin
	}
	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.Phone != "" {
		data["phone"] = req.Phone
	}
	if req.ProfilePictureURL != "" {
		data["profile_picture_url"] = req.ProfilePictureURL
	}
	if req.Status != "" {
		data["status"] = req.Status
	}
	log.Println(data,req)

	query, args, err := u.db.Sq.Builder.Update(u.tableName).
    SetMap(data).
    Where(u.db.Sq.Equal("\"user_id\"", req.UserID)).
    Suffix(fmt.Sprintf("RETURNING %s",selectQuery())).
    ToSql()

	if err != nil {
		return nil,err
	}
	
	var userResponse entityuser.User

	err = u.db.QueryRow(ctx,query,args...).Scan(
		&userResponse.UserID,
		&userResponse.Name,
		&userResponse.Email,
		&userResponse.Phone,
		&userResponse.PasswordHash,
		&userResponse.Status,
		&userResponse.ProfilePictureURL,
		&userResponse.Language,
		&userResponse.CreatedAt,
		&userResponse.LastLogin,
	)
	if err != nil {
		return nil,err
	}

	return &userResponse,err
}