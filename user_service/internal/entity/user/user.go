package entityuser

import "time"

type (
	User struct {
		UserID            string    `json:"user_id"`
		Name              string    `json:"name"`
		Email             string    `json:"email"`
		Phone             string    `json:"phone"`
		PasswordHash      string    `json:"password_hash"`
		Status            string    `json:"status"`
		ProfilePictureURL string    `json:"profile_picture_url"`
		Language          string    `json:"language"`
		CreatedAt         time.Time    `json:"created_at"`
		LastLogin         time.Time `json:"last_login"`
	}

	UpdateUserRequest struct {
		UserID            string `json:"user_id"`
		Name              string `json:"name"`
		Phone             string `json:"phone"`
		Status            string `json:"status"`
		ProfilePictureURL string `json:"profile_picture_url"`
		Language          string `json:"language"`
		LastLogin  string `json:"last_login"`
	}

	CreateUserRequest struct {
		Name              string `json:"name"`
		Email             string `json:"email"`
		Phone             string `json:"phone"`
		PasswordHash      string `json:"password_hash"`
		Status            string `json:"status"`
		ProfilePictureURL string `json:"profile_picture_url"`
		Language          string `json:"language"`
	}

	UserResponse struct {
		User User `json:"user"`
	}

	LoginUserRequest struct {
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
	}

	LoginUserResponse struct {
		Success bool `json:"success"`
	}

	GetUserRequest struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}

	GetAllUserRequest struct {
		Limit int64  `json:"limit"`
		Page  int64  `json:"page"`
		Field string `json:"field"`
		Value string `json:"value"`
	}

	GetUserAllResponse struct {
		Users []User `json:"users"`
	}

	ChangeUserPasswordRequest struct {
		UserID          string `json:"user_id"`
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	ChangeUserPasswordResponse struct {
		Success bool `json:"success"`
	}
)
