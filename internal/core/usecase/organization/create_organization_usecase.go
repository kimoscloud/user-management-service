package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	"github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	"github.com/kimoscloud/value-types/errors"
)

type CreateOrganizationUseCase struct {
	organizationRepository organization.Repository
	logger                 logging.Logger
}

func NewCreateOrganizationUseCase(
	organizationRepository organization.Repository,
	logger logging.Logger,
) CreateOrganizationUseCase {
	return CreateOrganizationUseCase{
		organizationRepository: organizationRepository,
		logger:                 logger,
	}
}

func (cu CreateOrganizationUseCase) Hadler(userId string, request CreateOrganizationRequest) (*CreateOrganizationResponse, *errors.AppError) {
	organization, err := cu.organizationRepository.Create()
	if err != nil {
		cu.logger.Error(err)
		return nil, &CreateOrganizationError{
			HTTPStatus: 500,
			Message:    "Internal server error",
		}
	}
	return &CreateOrganizationResponse{
		Id:   organization.Id,
		Name: organization.Name,
	}, nil
}
