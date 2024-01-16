package organization

type CreateOrganizationRequest struct {
	Name         string `json:"name"`
	BillingEmail string `json:"billingEmail" binding:"required"`
	Plan         string `json:"plan" binding:"required"`
	Captcha      string `json:"captcha" binding:"required"`
}
