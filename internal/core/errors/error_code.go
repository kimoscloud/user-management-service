package errors

import "github.com/kimoscloud/value-types/errors"

const ErrorUserAuthenticatedNotFound errors.ErrorCode = "0000007"
const ErrorUserEmailAlreadyExists errors.ErrorCode = "0000008"
const ErrorUserCantAddUsersToOrganization errors.ErrorCode = "0000009"
const ErrorUserCantRemoveUsersFromOrganization errors.ErrorCode = "00000020"
const ErrorUserCantCreateTeamIntoOrganization errors.ErrorCode = "00000021"
const ErrorTryingToGetTeamsByNameAndSlug errors.ErrorCode = "00000022"
const ErrorConflictTeamExistWithSameNameOrSlug errors.ErrorCode = "00000023"
const ErrorCreatingTeam errors.ErrorCode = "00000024"
const ErrorUserDontHavePrivilegesToAddTeamMembersToTeam errors.ErrorCode = "00000025"
