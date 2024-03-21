package controller

import (
	"github.com/gin-gonic/gin"
	auth2 "github.com/kimoscloud/user-management-service/internal/core/auth"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/model/request/auth"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/usecase"
	"github.com/kimoscloud/user-management-service/internal/middleware"
	"net/http"
	"strings"
)

type UserController struct {
	gin                      *gin.Engine
	createUserUseCase        *usecase.CreateUserUseCase
	authenticateUserUseCase  *usecase.AuthenticateUserUseCase
	getUserUseCase           *usecase.GetUserUseCase
	updateUserProfileUseCase *usecase.UpdateUserProfileUseCase
	changePasswordUseCase    *usecase.ChangePasswordUseCase
	logger                   logging.Logger
}

func NewUserController(
	gin *gin.Engine,
	logger logging.Logger,
	createUserUseCase *usecase.CreateUserUseCase,
	authenticateUserUseCase *usecase.AuthenticateUserUseCase,
	getUserUseCase *usecase.GetUserUseCase,
	updateUserProfileUseCase *usecase.UpdateUserProfileUseCase,
	changePasswordUseCase *usecase.ChangePasswordUseCase,
) UserController {
	return UserController{
		gin:                      gin,
		logger:                   logger,
		createUserUseCase:        createUserUseCase,
		authenticateUserUseCase:  authenticateUserUseCase,
		getUserUseCase:           getUserUseCase,
		updateUserProfileUseCase: updateUserProfileUseCase,
		changePasswordUseCase:    changePasswordUseCase,
	}
}

func (u UserController) InitRouter() {
	api := u.gin.Group("/api/v1/auth")
	api.POST("/signup", u.signUp)
	api.POST("/login", u.login)
	api.POST("/validate-token", u.validateToken)
	secured := api.Group("", middleware.Auth())
	{
		secured.GET("/me", u.me)
		secured.PUT("/me", u.updateProfile)
		secured.POST("/password", u.changePassword)
	}
}

func (u UserController) login(c *gin.Context) {
	signInRequest, err := u.parseLoginRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, &gin.H{
				"message": "Invalid request",
			},
		)
		return
	}
	result, appError := u.authenticateUserUseCase.Handler(*signInRequest)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

// validateController validates the user controller by retrieving the user ID from the request context
// and returning it in the JSON response.
func (u UserController) validateToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(
			401, gin.H{
				"message": "Unauthorized",
			},
		)
		c.Abort()
		return
	}
	authorizationHeaderSplitted := strings.Split(tokenString, "Bearer ")
	if len(authorizationHeaderSplitted) != 2 {
		c.AbortWithStatusJSON(
			401, gin.H{
				"message": "Invalid token",
			},
		)
		c.Abort()
		return
	}
	claims, err := auth2.ValidateToken(
		authorizationHeaderSplitted[1],
	)
	if err != nil {
		c.AbortWithStatusJSON(
			401, gin.H{
				"message": "Unauthorized",
			},
		)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, claims)
}

func (u UserController) changePassword(c *gin.Context) {
	userId := c.GetString("kimosUserId")
	changePasswordRequest, err := u.parseChangePasswordRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, &gin.H{
				"message": "Invalid request",
			},
		)
		return
	}
	changePasswordRequest.IsValid()
	appError := u.changePasswordUseCase.Handle(userId, changePasswordRequest)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
		},
	)
}

func (u UserController) signUp(c *gin.Context) {
	signUpRequest, err := u.parseSignUpRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, &gin.H{
				"message": "Invalid request",
			},
		)
		return
	}
	_, appError := u.createUserUseCase.Handler(signUpRequest)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.Status(http.StatusCreated)
	return
}

func (u UserController) updateProfile(c *gin.Context) {
	userId := c.GetString("kimosUserId")
	updateProfileRequest, err := u.parseUpdateProfileRequest(c)
	if err != nil {
		//TODO wrapp error
		c.AbortWithStatusJSON(
			http.StatusBadRequest, &gin.H{
				"message": "Invalid request",
			},
		)
		return
	}
	result, appError := u.updateUserProfileUseCase.Handler(
		userId,
		updateProfileRequest,
	)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u UserController) me(c *gin.Context) {
	userId := c.GetString("kimosUserId")
	result, appError := u.getUserUseCase.Handler(userId)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u UserController) parseSignUpRequest(ctx *gin.Context) (
	*auth.SignUpRequest,
	error,
) {
	var req auth.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (u UserController) parseLoginRequest(ctx *gin.Context) (
	*auth.LoginRequest,
	error,
) {
	var req auth.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (u UserController) parseUpdateProfileRequest(ctx *gin.Context) (
	*auth.UpdateProfileRequest,
	error,
) {
	var req auth.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (u UserController) parseChangePasswordRequest(c *gin.Context) (
	*request.ChangePasswordRequest,
	error,
) {
	var req request.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
