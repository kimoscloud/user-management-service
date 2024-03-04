package organization

import (
	errors2 "github.com/kimoscloud/user-management-service/internal/core/errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/constants"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	organizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	userOrganizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/user-organization"
	"github.com/kimoscloud/value-types/domain"
	"github.com/kimoscloud/value-types/errors"
)

type GetOrganizationMembersPagedUseCase struct {
	organizationRepo                    organizationRepository.Repository
	userOrgRepo                         userOrganizationRepository.Repository
	checkUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction
	logger                              logging.Logger
}

func NewGetOrganizationMembersPagedUseCase(
	organizationRepo organizationRepository.Repository,
	userOrgRepo userOrganizationRepository.Repository,
	checkUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction,
	logger logging.Logger,
) *GetOrganizationMembersPagedUseCase {
	return &GetOrganizationMembersPagedUseCase{
		organizationRepo:                    organizationRepo,
		userOrgRepo:                         userOrgRepo,
		checkUserHasPermissionsToMakeAction: checkUserHasPermissionsToMakeAction,
		logger:                              logger,
	}
}

func (cu GetOrganizationMembersPagedUseCase) Handler(
	authenticatedUserId, orgId, search string,
	pageNumber, pageSize int,
) (*domain.Page[organization.UserOrganization], *errors.AppError) {
	if !cu.checkUserHasPermissionsToMakeAction.Handler(
		authenticatedUserId,
		orgId,
		[]string{constants.PERMISSION_READ_ORGANIZATION_MEMBERS},
	) {
		return nil, errors2.NewForbiddenError(
			"The user don't have the privileges to do this operation",
			"The user don't have the privileges to do this operation if the error persist, contact with your administrator or contact us",
			errors2.ErrorUserDontHavePrivilegesToReadOrganizationMembers,
		).AppError
	}
	userOrganizations, err := cu.userOrgRepo.GetOrganizationMembersPaged(
		orgId, search, pageNumber,
		pageSize,
	)
	if err != nil {
		cu.logger.Error("Error getting organization members", "error", err.Error())
		return nil, errors.NewInternalServerError(
			"Error getting organization members",
			"",
			errors2.ErrorGettingOrganizationMembers,
		).AppError
	}
	return &userOrganizations, nil
}
