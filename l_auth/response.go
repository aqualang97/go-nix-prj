package main

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	NewAccessToken  string `json:"access_token"`
	NewRefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID    int
	Email string
	Name  string
}
