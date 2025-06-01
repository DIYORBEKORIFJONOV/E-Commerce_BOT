package middleware

import (
	"github.com/casbin/casbin/v2"
	tokensettings "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/token"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	rateLimiters = map[string]*rate.Limiter{
		"user":         rate.NewLimiter(5.0/60.0, 1),
		"unauthorized": rate.NewLimiter(5.0/60.0, 1),
	}
)

func Middleware(c *gin.Context) {
	allow, err := CheckPermission(c.Request)
	if err != nil {
		if valid, ok := err.(*jwt.ValidationError); ok && valid.Errors == jwt.ValidationErrorExpired {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "token was expired",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "permission denied",
		})
		return
	} else if !allow {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "permission denied",
		})
		return
	}
	role, _ := GetRole(c.Request)
	limiter, exists := rateLimiters[role]
	if !exists {
		limiter = rateLimiters["unauthorized"]
	}

	if !limiter.Allow() {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"error": "rate limit exceeded",
		})
		return
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	id, _ := tokensettings.GetIdFromToken(c.Request)
	c.Set("user_id", id)
	email, _ := tokensettings.GetEmailFromToken(c.Request)
	c.Set("email", email)
	c.Next()
}

func TimingMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	c.Writer.Header().Set("X-Response-Time", duration.String())
}

func CheckPermission(r *http.Request) (bool, error) {
	role, err := GetRole(r)
	if err != nil {
		log.Println("[DEBUG] Error getting role:", err)
	}
	log.Printf("[DEBUG] Casbin check → role=%s path=%s method=%s", role, r.URL.Path, r.Method)

	enforcer, err := casbin.NewEnforcer("auth.conf", "auth.csv")
	if err != nil {
		log.Println("[DEBUG] Enforcer error:", err)
		return false, err
	}

	allowed, err := enforcer.Enforce(role, r.URL.Path, r.Method)
	if err != nil {
		log.Println("[DEBUG] Enforce error:", err)
		return false, err
	}

	log.Printf("[DEBUG] Enforce result: allowed=%v", allowed)
	return allowed, nil
}

func GetRole(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		log.Println("[DEBUG] Missing or invalid Authorization header")
		return "unauthorized", nil
	}

	// Удаляем префикс "Bearer "
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := tokensettings.ExtractClaim(tokenStr)
	if err != nil {
		log.Println("[DEBUG] Error while extracting claims:", err)
		return "unauthorized", err
	}

	roleVal, ok := claims["role"]
	if !ok {
		log.Println("[DEBUG] role not found in claims")
		return "unauthorized", nil
	}

	roleStr, ok := roleVal.(string)
	if !ok {
		log.Println("[DEBUG] role claim is not a string")
		return "unauthorized", nil
	}

	log.Printf("[DEBUG] Extracted role: %s", roleStr)
	return roleStr, nil
}
