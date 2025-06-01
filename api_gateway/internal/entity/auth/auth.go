package authentity

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// RegisterForm содержит данные для регистрации пользователя
type RegisterForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
}

type Account struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password"`
	Phone        string `json:"phone"`
	Username     string `json:"username"`
	Role         string `json:"-"`
}

func (r *RegisterForm) Validate() (err error) {
	fields := map[string]string{
		"name":     r.Name,
		"password": r.Password,
		"phone":    r.Phone,
		"username": r.Username,
	}
	for key, val := range fields {
		val = strings.TrimSpace(val)

		if val == "" {
			return fmt.Errorf("%s is required", key)
		}

		switch key {
		case "phone":
			if !isValidUzbekPhone(val) {
				return errors.New("phone must be a valid Uzbekistan number (e.g., +998991234567)")
			}
		}
	}
	return
}

func isValidUzbekPhone(phone string) bool {
	re := regexp.MustCompile(`^(?:\+998|998)[0-9]{9}$`)
	return re.MatchString(phone)
}

// RegisterResponse представляет результат регистрации
type RegisterResponse struct {
	AlreadyExists bool `json:"already_exists"`
	Sent          bool `json:"sent"`
}

type VerifyRequest struct {
	Phone      string `json:"phone"`
	SecretCode string `json:"secret"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password"`
	Id          string `json:"-"`
}
