package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/usecase/user"
	"github.com/kimoscloud/user-management-service/internal/middleware"
	"net/http"
)

type UserController struct {
	gin                      *gin.Engine
	createUserUseCase        *user.CreateUserUseCase
	authenticateUserUseCase  *user.AuthenticateUserUseCase
	getUserUseCase           *user.GetUserUseCase
	updateUserProfileUseCase *user.UpdateUserProfileUseCase
	logger                   logging.Logger
}

func NewUserController(
	gin *gin.Engine,
	logger logging.Logger,
	createUserUseCase *user.CreateUserUseCase,
	authenticateUserUseCase *user.AuthenticateUserUseCase,
	getUserUseCase *user.GetUserUseCase,
	updateUserProfileUseCase *user.UpdateUserProfileUseCase,
) UserController {
	return UserController{
		gin:                      gin,
		logger:                   logger,
		createUserUseCase:        createUserUseCase,
		authenticateUserUseCase:  authenticateUserUseCase,
		getUserUseCase:           getUserUseCase,
		updateUserProfileUseCase: updateUserProfileUseCase,
	}
}

func (u UserController) InitRouter() {
	api := u.gin.Group("/api/v1/user")
	api.POST("/signup", u.signUp)
	api.POST("/login", u.login)
	secured := api.Group("", middleware.Auth())
	{
		secured.GET("/me", u.me)
		secured.PUT("/me", u.updateProfile)
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
	c.JSON(http.StatusCreated, gin.H{})
	return
}

func (u UserController) updateProfile(c *gin.Context) {
	userId := c.GetString("kimosUserId")
	updateProfileRequest, err := u.parseUpdateProfileRequest(c)
	if err != nil {
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
	*request.SignUpRequest,
	error,
) {
	var req request.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (u UserController) parseLoginRequest(ctx *gin.Context) (
	*request.LoginRequest,
	error,
) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (u UserController) parseUpdateProfileRequest(ctx *gin.Context) (
	*request.UpdateProfileRequest,
	error,
) {
	var req request.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
