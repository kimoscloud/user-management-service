package organization

import (
	errors2 "github.com/kimoscloud/user-management-service/internal/core/errors"
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
	organizationRepository              repository.Repository
	userOrganizationRepo                userOrganizationRepository.Repository
	roleRepo                            roleRepository.Repository
	userRepo                            userRepository.Repository
	checkUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction
	logger                              logging.Logger
}

func NewCreateOrganizationMemberUseCase(
	organizationRepository repository.Repository,
	userOrganizationRepo userOrganizationRepository.Repository,
	roleRepo roleRepository.Repository,
	userRepo userRepository.Repository,
	checkUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction,
	logger logging.Logger,
) *CreateOrganizationMemberUseCase {
	return &CreateOrganizationMemberUseCase{
		organizationRepository:              organizationRepository,
		userOrganizationRepo:                userOrganizationRepo,
		userRepo:                            userRepo,
		roleRepo:                            roleRepo,
		checkUserHasPermissionsToMakeAction: checkUserHasPermissionsToMakeAction,
		logger:                              logger,
	}
}

func (cu CreateOrganizationMemberUseCase) Handler(
	authenticatedUserId, orgId string,
	request *request.CreateOrganizationUsers,
) *errors.AppError {
	tx := cu.userOrganizationRepo.BeginTransaction()
	defer tx.Rollback()
	if !cu.checkUserHasPermissionsToMakeAction.Handler(authenticatedUserId, orgId, []string{domain.PERMISSION_ADD_ORGANIZATION_MEMBER}) {
		return errors2.NewForbiddenError("The user don't have the privileges to do this operation",
			"The user don't have the privileges to do this operation if the error persist, contact with your administrator or contact us support@kimos.cloud", "0000019").AppError
	}
	role, err := cu.roleRepo.GetRoleByIdAndOrgId(request.RoleId, orgId)
	if err != nil {
		cu.logger.Error("Error getting role by id and org id", err)
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error inviting user to the organization",
			"Error searching the role",
			"0000012",
		).AppError
	}
	if role == nil {
		cu.logger.Error("Error: role was not found", err)
		//TODO replace the error code here
		return errors.NewNotFoundError(
			"Error inviting user to the organization",
			"Role not found",
			"0000013",
		).AppError
	}
	users, err := cu.userRepo.FindUsersByEmails(request.Emails)
	if err != nil {
		cu.logger.Error("Error inviting user to the organization", err)
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error inviting user to the organization",
			"Error searching the users",
			"0000014",
		).AppError
	}
	var userOrganizations []organization.UserOrganization
	var userIds []string
	for _, user := range users {
		userIds = append(userIds, user.ID)
		userOrg := organization.UserOrganization{
			UserID:         user.ID,
			OrganizationID: orgId,
			RoleID:         role.ID,
			InvitedAt:      time.Now(),
		}
		userOrganizations = append(userOrganizations, userOrg)
	}
	orgUsers, err := cu.userOrganizationRepo.
		GetUserOrganizationsByUserIdsAndOrganizationIdIgnoreDeletedAt(
			userIds, orgId,
		)
	if err != nil {
		cu.logger.Error("Error searching the user organizations", err)
		//TODO replace the error code here
		return errors.NewInternalServerError(
			"Error inviting user to the organization",
			"Error searching the user organizations",
			"0000015",
		).AppError
	}
	var userOrganizationIdsToBeRestored []string
	//Filter users that exists before create
	for _, userOrg := range orgUsers {
		for i, userOrganization := range userOrganizations {
			if userOrg.UserID == userOrganization.UserID {
				userOrganizations = append(userOrganizations[:i], userOrganizations[i+1:]...)
				break
			}
		}
		if userOrg.DeletedAt.Valid {
			userOrganizationIdsToBeRestored = append(userOrganizationIdsToBeRestored, userOrg.ID)
		}
	}
	if len(userOrganizations) > 0 {
		err = cu.userOrganizationRepo.CreateUserOrganizations(userOrganizations, tx)
		if err != nil {
			tx.Rollback()
			cu.logger.Error("Error inviting user to the organization", err)
			//TODO replace the error code here
			return errors.NewInternalServerError(
				"Error inviting user to the organization",
				"Error creating user organizations",
				"0000015",
			).AppError
		}
	}

	if len(userOrganizationIdsToBeRestored) > 0 {
		err = cu.userOrganizationRepo.RestoreUserOrganizations(userOrganizationIdsToBeRestored, tx)
		if err != nil {
			tx.Rollback()
			cu.logger.Error("Error inviting user to the organization", err)
			//TODO replace the error code here
			return errors.NewInternalServerError(
				"Error inviting user to the organization",
				"Error restoring user organizations",
				"0000016",
			).AppError
		}
	}
	//TODO send email here to new users
	tx.Commit()
	return nil
}
