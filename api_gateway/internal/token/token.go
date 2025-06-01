package tokensettings

import (
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/core"
	"github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/config"
	authentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/auth"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"
)

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Token()), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return claims, nil
}

func GetIdFromToken(r *http.Request) (string, int) {
	var softToken string
	token := r.Header.Get("Authorization")

	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	} else if strings.Contains(token, "Bearer") {
		softToken = strings.TrimPrefix(token, "Bearer ")
	} else {
		softToken = token
	}

	claims, err := ExtractClaim(softToken)
	if err != nil {
		return "unauthorized", http.StatusUnauthorized
	}

	return cast.ToString(claims["uid"]), 0
}

func GetEmailFromToken(r *http.Request) (string, int) {
	var softToken string
	token := r.Header.Get("Authorization")

	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	} else if strings.Contains(token, "Bearer") {
		softToken = strings.TrimPrefix(token, "Bearer ")
	} else {
		softToken = token
	}

	claims, err := ExtractClaim(softToken)
	if err != nil {
		return "unauthorized", http.StatusUnauthorized
	}

	return cast.ToString(claims["email"]), 0
}

var jwtSecret = []byte("3a1G90F5734gj5o513F45lh")

func GenerateToken(account authentity.Account) (string, error) {
	claims := jwt.MapClaims{
		"id":       account.Id,
		"name":     account.Name,
		"username": account.Username,
		"phone":    account.Phone,
		"role":     "user",
		"exp":      time.Now().Add(core.TOKENEXPIRE_TIME).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
