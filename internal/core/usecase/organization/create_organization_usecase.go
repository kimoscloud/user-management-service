package organization

import (
	"github.com/kimoscloud/user-management-service/internal/core/model/entity/organization"
	request "github.com/kimoscloud/user-management-service/internal/core/model/response/organization"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	repository "github.com/kimoscloud/user-management-service/internal/core/ports/repository/organization"
	"github.com/kimoscloud/value-types/errors"
)

type CreateOrganizationUseCase struct {
	organizationRepository repository.Repository
	logger                 logging.Logger
}

func NewCreateOrganizationUseCase(
	organizationRepository repository.Repository,
	logger logging.Logger,
) CreateOrganizationUseCase {
	return CreateOrganizationUseCase{
		organizationRepository: organizationRepository,
		logger:                 logger,
	}
}

func (cu CreateOrganizationUseCase) Handler(userId string, request request.CreateOrganizationRequest) (*organization.Organization, *errors.AppError) {
	organizationResult, err := cu.organizationRepository.Create(&organization.Organization{
		Name:      request.Name,
		BillingEmail: request.BillingEmail
		CreatedBy: userId,
	})
	if err != nil {
		return nil, errors.NewInternalServerError(
			"Error getting user by email",
			"",
			errors.ErrorCreatingUser,
		).AppError
	}
	return organizationResult, nil
}
