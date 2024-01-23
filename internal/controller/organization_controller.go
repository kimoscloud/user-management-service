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
	gin                             *gin.Engine
	createOrganizationUseCase       *usecase.CreateOrganizationUseCase
	getOrganizationsByUserIdUseCase *usecase.GetOrganizationsByUserUseCase
	logger                          logging.Logger
}

func NewOrganizationController(
	gin *gin.Engine,
	logger logging.Logger,
	createOrganizationUseCase *usecase.CreateOrganizationUseCase,
	getOrganizationsByUserIdUseCase *usecase.GetOrganizationsByUserUseCase,
) OrganizationController {
	return OrganizationController{
		gin:                             gin,
		logger:                          logger,
		createOrganizationUseCase:       createOrganizationUseCase,
		getOrganizationsByUserIdUseCase: getOrganizationsByUserIdUseCase,
	}
}

func (oc OrganizationController) InitRouter() {
	api := oc.gin.Group("/api/v1/organization", middleware.Auth())
	api.GET("/:orgId/teams/:teamId/members", oc.getTeamMembers)
	api.POST("/:orgId/teams/:teamId/members", oc.addTeamMembers)
	api.DELETE("/:orgId/teams/:teamId/members", oc.removeTeamMembers)
	//api.POST("/:orgId/teams/:teamId/clone", oc.cloneTeam)
	api.GET("/:orgId/teams/:teamId", oc.getTeamById)
	api.PUT("/:orgId/teams/:teamId", oc.updateTeam)
	api.DELETE("/:orgId/teams/:teamId", oc.deleteTeam)
	api.GET("/:orgId/teams", oc.getTeamsByOrganizationIdAndUser)
	api.POST("/:orgId/teams", oc.createTeam)
	api.GET("/:orgId/members/:memberId", oc.getOrganizationMemberById)
	api.DELETE("/:orgId/members/:memberId", oc.removeOrganizationMember)
	api.POST("/:orgId/members", oc.createOrganizationMember)
	//TODO implement billing methods
	api.PUT("/:orgId", oc.updateOrganization)
	api.DELETE("/:orgId", oc.deleteOrganization)
	api.POST("", oc.createOrganization)
	api.GET("", oc.getOrganizations)
}

func (oc OrganizationController) getTeamMembers(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) addTeamMembers(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) removeTeamMembers(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) getTeamById(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) updateTeam(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) deleteTeam(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) getTeamsByOrganizationIdAndUser(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) createTeam(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) getOrganizationMemberById(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) removeOrganizationMember(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) createOrganizationMember(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) updateOrganization(c *gin.Context) {
	//TODO implement
}

func (oc OrganizationController) deleteOrganization(c *gin.Context) {
	//TODO implement desactive organization and left a job to be runned
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
	result, appError := oc.createOrganizationUseCase.Handler(userId, organization)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (oc OrganizationController) getOrganizations(c *gin.Context) {
	userId := c.GetString("kimosUserId")
	result, appError := oc.getOrganizationsByUserIdUseCase.Handler(userId)
	if appError != nil {
		c.AbortWithStatusJSON(appError.HTTPStatus, appError)
		return
	}
	c.JSON(http.StatusOK, result)
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
