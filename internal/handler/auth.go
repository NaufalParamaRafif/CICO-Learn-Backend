package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NaufalParamaRafif/CICO-Learn-Backend/domain"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/dto/request"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/dto/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUseCase domain.UserUseCase
}

func NewAuthHandler(
	router *gin.Engine,
	userUseCase domain.UserUseCase,
	// Auth midd
) {
	h := AuthHandler{
		userUseCase: userUseCase,
	}

	// Auth
	router.POST("/auth/signup", h.Signup)
	router.POST("/auth/login", h.Login)
	// router.POST("/auth/logout", h.Logout)
	// router.POST("/refresh")
	// router.POST("/validate")
}

func (h AuthHandler) Login(ctx *gin.Context) {
	loginInput := request.LoginRequest{}

	if err := ctx.Bind(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		fmt.Println("Error: " + err.Error())

		return
	}
	gin.DisableBindValidation()

	loginResult, err := h.userUseCase.Login(ctx, domain.LoginInput{
		Identifier: loginInput.Identifier,
		Password:   loginInput.Password,
	})
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrUserNotFound.Error(),
			})
		} else if errors.Is(err, domain.ErrInvalidIdentifierPassword) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": domain.ErrInvalidIdentifierPassword.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			fmt.Println("Error: " + err.Error())
		}

		return
	}

	ctx.JSON(http.StatusOK, response.LoginResponse{
		RefreshToken: loginResult.RefreshToken,
		AccessToken:  loginResult.AccessToken,
	})
}

func (h AuthHandler) Signup(ctx *gin.Context) {
	var req request.UserRegisterRequest

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	gin.DisableBindValidation()

	if err := h.userUseCase.Register(ctx, domain.RegisterInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		if errors.Is(err, domain.ErrEmailExists) {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": domain.ErrEmailExists.Error(),
			})
		} else if errors.Is(err, domain.ErrUsernameTaken) {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": domain.ErrUsernameTaken.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			fmt.Println("Error: " + err.Error())
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signup success",
	})
}

// func (h AuthHandler) Logout(ctx *gin.Context) {
// TODO: Revoke = true reddis
// }
