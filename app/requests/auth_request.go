package requests

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"customer@app.com"`
	Password string `json:"password" validate:"required" example:"password" `
}

type AuthRegisterRequest struct {
	Name     string `json:"name" validate:"required" example:"Jhon" `
	Email    string `json:"email" validate:"required,email" example:"customer@app.com"`
	Password string `json:"password" validate:"required" example:"password" `
}

type AuthForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email" example:"customer@app.com"`
}

type AuthResetPasswordRequest struct {
	Email       string `json:"email" validate:"required,email" example:"customer@app.com"`
	CodeToken   string `json:"code_token" validate:"required" example:"customer@app.com"`
	NewPassword string `json:"new_password" validate:"required" example:"customer@app.com"`
}
