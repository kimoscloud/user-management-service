package request

import "github.com/kimoscloud/value-types/is_valid"

type ChangePasswordRequest struct {
	OldPassword        string `json:"oldPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

func (r ChangePasswordRequest) IsValid() bool {
	if r.OldPassword == "" {
		return false
	}
	if r.NewPassword == "" || !is_valid.IsValidPassword(r.NewPassword) {
		return false
	}
	if r.NewPassword != r.ConfirmNewPassword {
		return false
	}
	return true
}
