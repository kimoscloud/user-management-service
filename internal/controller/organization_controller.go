package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/usecase/organization"
)

type OrganizationController struct {
	gin                       *gin.Engine
	createOrganizationUseCase *organization.CreateOrganizationUseCase
	logger                    logging.Logger
}
