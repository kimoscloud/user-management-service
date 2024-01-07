package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
)

type CreateOrganizationUseCase struct {
	organizationRepository organization.Repository
	logger                 logging.Logger
}
