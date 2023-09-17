package api

var (
	ClientID     = "YakK2QgXMECO2j5kmaHY"
	ClientSecret = "27spzmSVPFQ7s2eL9shZD9Gv9k5G45MnjOjsACaT"
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
