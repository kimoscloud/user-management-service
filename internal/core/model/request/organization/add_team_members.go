package organization

type AddTeamMembersRequest struct {
	UserIds []string `json:"userIds"`
	Role    string   `json:"role"`
}
