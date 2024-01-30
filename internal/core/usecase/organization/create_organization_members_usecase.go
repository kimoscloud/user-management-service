package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	request "github.com/kimoscloud/user-management-service/internal/core/model/request/organization"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	repository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	roleRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/role"
	userOrganizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/user-organization"
	userRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/user"
	"github.com/kimoscloud/value-types/domain"
	"github.com/kimoscloud/value-types/errors"
	"time"
)

type CreateOrganizationMemberUseCase struct {
	organizationRepository repository.Repository
	userOrganizationRepo   userOrganizationRepository.Repository
	roleRepo               roleRepository.Repository
	userRepo               userRepository.Repository
	logger                 logging.Logger
}

func NewCreateOrganizationMemberUseCase(
	organizationRepository repository.Repository,
	userOrganizationRepo userOrganizationRepository.Repository,
	roleRepo roleRepository.Repository,
	logger logging.Logger,
) *CreateOrganizationMemberUseCase {
	return &CreateOrganizationMemberUseCase{
		organizationRepository: organizationRepository,
		userOrganizationRepo:   userOrganizationRepo,
		roleRepo:               roleRepo,
		logger:                 logger,
	}
}

func (cu CreateOrganizationMemberUseCase) Handler(
	authenticatedUserId, orgId string,
	request *request.CreateOrganizationUsers,
) *errors.AppError {
	authenticatedOrgUser, err := cu.userOrganizationRepo.GetUserOrganizationByUserAndOrganizationWithRolesAndPermissions(
		authenticatedUserId,
		orgId,
	)
	if err != nil {
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error inviting user to the organization",
			"Error searching the authenticated user",
			"0000011",
		).AppError
	}
	authenticatedOrgUser.CheckIfOrgUserHasPermissions(
		[]string{domain.PERMISSION_ADD_ORGANIZATION_MEMBER},
	)
	role, err := cu.roleRepo.GetRoleByIdAndOrgId(request.RoleId, orgId)
	if err != nil {
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error inviting user to the organization",
			"Error searching the role",
			"0000012",
		).AppError
	}
	if role == nil {
		//TODO replace the error code here
		return errors.NewNotFoundError(
			"Error inviting user to the organization",
			"Role not found",
			"0000013",
		).AppError
	}
	users, err := cu.userRepo.FindUsersByEmails(request.Emails)
	if err != nil {
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error inviting user to the organization",
			"Error searching the users",
			"0000014",
		).AppError
	}
	var userOrganizations []organization.UserOrganization
	for _, user := range users {
		userOrg := organization.UserOrganization{
			UserID:         user.ID,
			OrganizationID: orgId,
			RoleID:         role.ID,
			InvitedAt:      time.Now(),
		}
		userOrganizations = append(userOrganizations, userOrg)
	}
	err = cu.userOrganizationRepo.CreateUserOrganizations(userOrganizations, nil)

	return nil
}
