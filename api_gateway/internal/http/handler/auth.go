package handler

import (
	authentity "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/auth"
	models "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/entity/order"
	authusecase "github.com/diyorbek/E-Commerce_BOT/api_gateway/internal/usecase/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthHandler struct {
	auth authusecase.AuthUseCaseIml
}

func NewAuthHandler(auth authusecase.AuthUseCaseIml) *AuthHandler {
	return &AuthHandler{
		auth: auth,
	}
}

// @title Artisan Connect
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host hurmomarketshop.duckdns.org
// @BasePath        /
// @schemes         https
// @securityDefinitions.apiKey ApiKeyAuth
// @in              header
// @name            Authorization

// RegisterAccountHandler godoc
// @Summary      Retrieve all orders
// @Description  Регистрация нового пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body   authentity.RegisterForm  true  "Register form"
// @Success      200    {object} authentity.RegisterResponse "response"
// @Failure      400    {object} models.ErrorResponse    "Bad Request"
// @Failure      500    {object} models.ErrorResponse    "Internal Server Error"
// @Router       /auth/register [post]
func (h *AuthHandler) RegisterAccountHandler(ginContext *gin.Context) {
	var requestModel authentity.RegisterForm
	if err := ginContext.ShouldBind(&requestModel); err != nil {
		ginContext.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := requestModel.Validate(); err != nil {
		ginContext.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	switch result, err := h.auth.RegisterAccount(ginContext.Request.Context(), &requestModel); {
	case err != nil:
		ginContext.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	case result.AlreadyExists:
		ginContext.JSON(http.StatusConflict, models.ErrorResponse{Error: "account already exists"})
		return
	default:
		ginContext.JSON(http.StatusOK, result)
	}
}

// VerifyAccountHandler godoc
// @Summary      Verify user account
// @Description  Подтверждение аккаунта по коду (например, OTP)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body   authentity.VerifyRequest  true  "Verification request"
// @Success      200  {object}   string     "Успешная верификация"
// @Failure      400  {object}  models.ErrorResponse         "Ошибка валидации"
// @Failure      500  {object}  models.ErrorResponse         "Внутренняя ошибка сервера"
// @Router       /auth/verify [post]
func (h *AuthHandler) VerifyAccountHandler(ginContext *gin.Context) {
	var requestMode authentity.VerifyRequest
	if err := ginContext.ShouldBindJSON(&requestMode); err != nil {
		ginContext.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	token, err := h.auth.VerifyAccount(ginContext.Request.Context(), &requestMode)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	ginContext.JSON(http.StatusOK, token)
}

// LoginHandler godoc
// @Summary      Login user
// @Description  Авторизация пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body   authentity.LoginRequest  true  "Login request"
// @Success      200  {object}   string   "Успешный вход"
// @Failure      400  {object}  models.ErrorResponse       "Ошибка валидации"
// @Failure      500  {object}  models.ErrorResponse       "Внутренняя ошибка сервера"
// @Router       /auth/login [post]
func (h *AuthHandler) LoginHandler(ginContext *gin.Context) {
	var requestMode authentity.LoginRequest
	if err := ginContext.ShouldBind(&requestMode); err != nil {
		ginContext.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	token, err := h.auth.Login(ginContext.Request.Context(), requestMode.Username, requestMode.Password)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	ginContext.JSON(http.StatusOK, token)
}

// ChangePasswordHandler godoc
// @Summary      Login user
// @Description изменит пароль пользователя
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        payload  body   authentity.ChangePasswordRequest  true  "Login request"
// @Security ApiKeyAuth
// @Success      200  {object}   string   "Успешный изменения"
// @Failure      400  {object}  models.ErrorResponse       "Ошибка валидации"
// @Failure      500  {object}  models.ErrorResponse       "Внутренняя ошибка сервера"
// @Router       /account/update-password [put]
func (h *AuthHandler) ChangePasswordHandler(ginContext *gin.Context) {
	id, ok := ginContext.Get("user_id")
	if !ok {
		log.Println(id, "ewgwegwe")
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	var requestModel authentity.ChangePasswordRequest
	if err := ginContext.ShouldBindJSON(&requestModel); err != nil {
		ginContext.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	response, err := h.auth.ChangePassword(ginContext.Request.Context(), id.(string), requestModel.NewPassword)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	ginContext.JSON(http.StatusOK, map[string]interface{}{
		"changed": response,
	})

}
