package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	repository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	"github.com/kimoscloud/value-types/errors"
)

type GetOrganizationsByUserUseCase struct {
	organizationRepository repository.Repository
	logger                 logging.Logger
}

func NewGetOrganizationsByUserUseCase(
	organizationRepository repository.Repository,
	logger logging.Logger,
) *GetOrganizationsByUserUseCase {
	return &GetOrganizationsByUserUseCase{
		organizationRepository: organizationRepository,

		logger: logger,
	}
}

func (cu GetOrganizationsByUserUseCase) Handler(
	userId string,
) ([]organization.Organization, *errors.AppError) {
	organizationResult, err := cu.organizationRepository.GetAllByUserId(userId)
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting user organizations",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	return organizationResult, nil
}
