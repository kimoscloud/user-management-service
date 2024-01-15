package controller

import (
	"github.com/gin-gonic/gin"
	organizationRequest "github.com/kimoscloud/user-management-service/internal/core/model/request/organization"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	usecase "github.com/kimoscloud/user-management-service/internal/core/usecase/organization"
	"github.com/kimoscloud/user-management-service/internal/middleware"
	"net/http"
)

type OrganizationController struct {
	gin                       *gin.Engine
	createOrganizationUseCase *usecase.CreateOrganizationUseCase
	logger                    logging.Logger
}

func NewOrganizationController(
	gin *gin.Engine,
	logger logging.Logger,
	createOrganizationUseCase *usecase.CreateOrganizationUseCase,
) OrganizationController {
	return OrganizationController{
		gin:                       gin,
		logger:                    logger,
		createOrganizationUseCase: createOrganizationUseCase,
	}
}

func (oc OrganizationController) InitRouter() {
	api := oc.gin.Group("/api/v1/organization", middleware.Auth())
	api.POST("", oc.createOrganization)
}

func (oc OrganizationController) createOrganization(c *gin.Context) {
	userId := c.GetString("kimosUserId")
	organization, err := oc.parseCreateOrganizationRequest(c)
	if err != nil {
		//TODO wrapp error
		c.AbortWithStatusJSON(
			400, &gin.H{
				"message": "Invalid request",
			},
		)
		return
	}
	result, appError := oc.createOrganizationUseCase.Handler(userId, *organization)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (oc OrganizationController) parseCreateOrganizationRequest(ctx *gin.Context) (
	*organizationRequest.CreateOrganizationRequest,
	error,
) {
	var request organizationRequest.CreateOrganizationRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}
