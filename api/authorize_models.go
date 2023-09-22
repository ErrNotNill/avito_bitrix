package api

import "os"

var (
	ClientID     = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	Token        = ""
)

type Auth struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type Notification struct {
	Secret string `json:"secret"`
	Url    string `json:"url"`
}
