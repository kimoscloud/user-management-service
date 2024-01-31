package organization

type CreateOrganizationUsers struct {
	Emails []string `json:"emails"`
	RoleId string   `json:"roleId"`
}
