package request

type SignUpRequest struct {
	Email                    string `json:"email"`
	Password                 string `json:"password"`
	ConfirmPassword          string `json:"confirmPassword"`
	AcceptTermsAndConditions bool   `json:"acceptTermsAndConditions"`
}
