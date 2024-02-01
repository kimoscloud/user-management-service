package organization

import (
	errors2 "github.com/kimoscloud/user-management-service/internal/core/errors"
	organizationRequest "github.com/kimoscloud/user-management-service/internal/core/model/request/organization"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	teamRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/team"
	teamMemberRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/team-member"
	userOrganizationRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization/user-organization"
	"github.com/kimoscloud/value-types/errors"
)

type AddTeamMembersUseCase struct {
	userOrganizationRepo                  userOrganizationRepository.Repository
	teamRepo                              teamRepository.Repository
	teamMemberRepo                        teamMemberRepository.Repository
	checkIfUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction
	logger                                logging.Logger
}

func NewAddTeamMembersUseCase(userOrganizationRepo userOrganizationRepository.Repository,
	teamRepo teamRepository.Repository,
	teamMemberRepo teamMemberRepository.Repository,
	checkIfUserHasPermissionsToMakeAction *CheckUserHasPermissionsToMakeAction,
	logger logging.Logger) *AddTeamMembersUseCase {
	return &AddTeamMembersUseCase{
		userOrganizationRepo:                  userOrganizationRepo,
		teamRepo:                              teamRepo,
		teamMemberRepo:                        teamMemberRepo,
		checkIfUserHasPermissionsToMakeAction: checkIfUserHasPermissionsToMakeAction,
		logger:                                logger,
	}
}

func (u *AddTeamMembersUseCase) Handler(authenticatedUserId string, orgId string, request *organizationRequest.AddTeamMembersRequest) *errors.AppError {
	tx := u.teamRepo.BeginTransaction()
	defer tx.Rollback()
	if !u.checkIfUserHasPermissionsToMakeAction.Handler(authenticatedUserId, orgId, []string{"ADD_TEAM_MEMBERS"}) {
		return errors2.NewForbiddenError("The user don't have the privileges to do this operation",
			"The user don't have the privileges to do this operation if the error persist, contact with your administrator or contact us",
			errors2.ErrorUserDontHavePrivilegesToAddTeamMembersToTeam).AppError
	}
	// I should be able to invite a new member form here  (To org too)
	return nil
}
