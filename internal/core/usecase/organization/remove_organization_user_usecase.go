package organization

import (
	errors2 "github.com/kimoscloud/user-management-service/internal/core/errors"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	repository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	userOrganizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/user-organization"
	"github.com/kimoscloud/value-types/domain"
	"github.com/kimoscloud/value-types/errors"
)

type RemoveOrganizationMemberUseCase struct {
	organizationRepository repository.Repository
	userOrganizationRepo   userOrganizationRepository.Repository
	logger                 logging.Logger
}

func NewRemoveOrganizationMemberUseCase(
	organizationRepository repository.Repository,
	userOrganizationRepo userOrganizationRepository.Repository,
	logger logging.Logger,
) *RemoveOrganizationMemberUseCase {
	return &RemoveOrganizationMemberUseCase{
		organizationRepository: organizationRepository,
		userOrganizationRepo:   userOrganizationRepo,
		logger:                 logger,
	}
}

func (cu RemoveOrganizationMemberUseCase) Handler(
	authenticatedUserId, orgId string,
	memberId string,
) *errors.AppError {
	tx := cu.userOrganizationRepo.BeginTransaction()
	defer tx.Rollback()
	authenticatedOrgUser, err := cu.userOrganizationRepo.GetUserOrganizationByUserAndOrganizationWithRolesAndPermissions(
		authenticatedUserId,
		orgId,
	)
	if err != nil {
		cu.logger.Error("Error searching the authenticated user", err)
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error removing user from the organization",
			"Error searching the authenticated user",
			"0000011",
		).AppError
	}
	if !authenticatedOrgUser.CheckIfOrgUserHasPermissions(
		[]string{domain.PERMISSION_REMOVE_ORGANIZATION_TEAM_MEMBER},
	) {
		return errors2.NewForbiddenError(
			"Error removing user from the organization",
			"Authenticated user does not have permission to remove user from the organization",
			errors2.ErrorUserCantRemoveUsersFromOrganization,
		).AppError
	}
	if err := cu.userOrganizationRepo.RemoveUserFromOrganization(
		memberId,
		orgId,
		nil,
	); err != nil {
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error removing user from the organization",
			"Error removing user from the organization",
			"0000011",
		).AppError
	}
	return nil
}
