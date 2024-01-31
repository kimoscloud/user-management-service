package organization

type CreateTeamRequest struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	About string `json:"about"`
}
