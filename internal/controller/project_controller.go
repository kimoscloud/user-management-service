package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
)

type ProjectController struct {
	gin    *gin.Engine
	logger logging.Logger
}

func NewProjectController(
	gin *gin.Engine,
	logger logging.Logger,
) ProjectController {
	return ProjectController{
		gin:    gin,
		logger: logger,
	}
}

func (pc ProjectController) InitRouter() {
	api := pc.gin.Group("/api/v1/project")
	//api.GET("/:projectId/applications/:applicationId", pc.getApplicationById)
	//api.PUT("/:projectId/applications/:applicationId", pc.getApplicationById)
	//api.DELETE("/:projectId/applications/:applicationId", pc.deleteApplicationById)
	//api.POST("/:projectId/applications", pc.createApplication)
	//api.POST("/:projectId/applications", pc.getApplicationsByProjectId) //create application
	//api.POST("/:projectId/teams", pc.assignTeamToApplication)           //
	//api.PUT("/:projectId/teams/:teamId", pc.assignTeamToApplication)
	//api.DELETE("/:projectId/teams", pc.assignTeamToApplication)
	api.GET("/:projectId", pc.getProjectById)
	api.PUT("/:projectId", pc.updateProject)
	api.DELETE("/:projectId", pc.deleteProject)
	api.GET("", pc.getProjects)
	api.POST("", pc.createProject)
}

func (pc ProjectController) getProjectById(c *gin.Context) {
	//TODO implement
}

func (pc ProjectController) updateProject(c *gin.Context) {
	//TODO implement
}

func (pc ProjectController) deleteProject(c *gin.Context) {
	//TODO implement
}

func (pc ProjectController) getProjects(c *gin.Context) {
	//TODO implement
}

func (pc ProjectController) createProject(c *gin.Context) {
	//TODO implement
}
