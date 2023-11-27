package models

type ResetPasswordReq struct {
	Email string `json:"email" validate:"required"`
	DOB   string `json:"dob" validate:"required"`
}

type NewPasswordReq struct {
	Otp             string `json:"otp" validate:"required"`
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
}
