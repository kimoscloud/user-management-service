package organization

type CreateOrganizationRequest struct {
	Name         string `json:"name" binding:"required"`
	BillingEmail string `json:"billingEmail" binding:"required"`
}
