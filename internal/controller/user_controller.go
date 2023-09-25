package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kimoscloud/user-management-service/internal/core/model/request"
	"github.com/kimoscloud/user-management-service/internal/core/usecase"
	"net/http"
)

type UserController struct {
	gin               *gin.Engine
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserController(
	gin *gin.Engine,
	createUserUseCase *usecase.CreateUserUseCase) UserController {
	return UserController{
		gin:               gin,
		createUserUseCase: createUserUseCase,
	}
}

func (u UserController) InitRouter() {
	api := u.gin.Group("/api/v1")
	api.POST("/signup", u.signUp)
}

func (u UserController) signUp(c *gin.Context) {
	_, err := u.parseSignUpRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &gin.H{
			"message": "Invalid request",
		})
		return
	}

	_, err = u.createUserUseCase.Handler(nil)
	if err != nil {
		return
	}
}

func (u UserController) parseSignUpRequest(ctx *gin.Context) (*request.SignUpRequest, error) {
	var req request.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
