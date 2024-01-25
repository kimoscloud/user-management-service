package organization

type OrganizationResponse struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Slug               string `json:"slug"`
	ImageUrl           string `json:"imageUrl"`
	LogoUrl            string `json:"logoUrl"`
	BackgroundImageUrl string `json:"backgroundImageUrl"`
	BillingEmail       string `json:"billingEmail"`
	Timezone           string `json:"timezone"`
	URL                string `json:"url"`
}
