package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	repository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	"github.com/kimoscloud/value-types/errors"
)

type GetOrganizationByOrgIAndUserIddUseCase struct {
	orgRepository repository.Repository
	logger        logging.Logger
}

func NewGetOrganizationByOrgIdAndUserIdUseCase(
	orgRepository repository.Repository,
	logger logging.Logger,
) *GetOrganizationByOrgIAndUserIddUseCase {
	return &GetOrganizationByOrgIAndUserIddUseCase{
		orgRepository: orgRepository,
		logger:        logger,
	}
}

func (cu GetOrganizationByOrgIAndUserIddUseCase) Handler(
	orgId string,
	userId string,
) (*organization.Organization, *errors.AppError) {
	org, err := cu.orgRepository.GetByIDAndUserId(orgId, userId)
	if err != nil {
		cu.logger.Error(err.Error())
		return nil, errors.NewInternalServerError(
			"error getting the organization",
			"an error occurred while getting the organization", "00001", //TODO update code
		).AppError
	}
	if org.ID == "" {
		return nil, errors.NewNotFoundError(
			"organization not found",
			"the organization with the given id was not found", "00002", //TODO update code
		).AppError
	}
	return org, nil
}
