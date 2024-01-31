package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	userOrganizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/user-organization"
)

type CheckUserHasPermissionsToMakeAction struct {
	userOrganizationRepo userOrganizationRepository.Repository
	logger               logging.Logger
}

func NewCheckUserHasPermissionsToMakeAction(userOrganizationRepo userOrganizationRepository.Repository,
	logger logging.Logger) *CheckUserHasPermissionsToMakeAction {
	return &CheckUserHasPermissionsToMakeAction{
		userOrganizationRepo: userOrganizationRepo,
		logger:               logger,
	}
}

func (oc *CheckUserHasPermissionsToMakeAction) Handler(userId, organizationId string, permissions []string) bool {
	authenticatedOrgUser, err := oc.userOrganizationRepo.GetUserOrganizationByUserAndOrganizationWithRolesAndPermissions(
		userId,
		organizationId,
	)
	if err != nil {
		oc.logger.Error("Error getting the organization user", err)
		return false
	}
	return authenticatedOrgUser.CheckIfOrgUserHasPermissions(
		permissions,
	)
}
