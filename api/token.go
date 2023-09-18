package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateAccessToken() string {
	ClientIDSettings := FindSettings(ClientID)
	ClientSecretSettings := FindSettings(ClientSecret)
	client := &http.Client{}
	auth := &Auth{}
	grantType := "client_credentials"
	url := fmt.Sprintf(`https://api.avito.ru/token?client_id=%s&client_secret=%s&grant_type=%s`, ClientIDSettings, ClientSecretSettings, grantType)
	post, err := client.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Println("Failed to create token: ", err)
	}
	body, err := io.ReadAll(post.Body)
	json.Unmarshal(body, &auth)
	fmt.Println("string(body).CreateAccessToken : ", string(body))
	fmt.Println("status code: ", post.StatusCode)
	fmt.Println("authorize.AccessToken: ", auth.AccessToken)
	token := GetToken()
	fmt.Println("token from DB: ", token)
	//AddToken(auth)
	return auth.AccessToken
}

func RefreshAccessToken() string {
	client := &http.Client{}
	auth := &Auth{}
	grantType := "refresh_token"
	url := fmt.Sprintf(`https://api.avito.ru/token?client_id=%s&client_secret=%s&grant_type=%s`, ClientID, ClientSecret, grantType)
	post, err := client.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Println("Failed to create token: ", err)
	}
	body, err := io.ReadAll(post.Body)
	json.Unmarshal(body, &auth)
	fmt.Println("string(body).CreateAccessToken : ", string(body))
	fmt.Println("status code: ", post.StatusCode)
	fmt.Println("authorize.AccessToken: ", auth.AccessToken)
	//Token = authorize.AccessToken
	return auth.AccessToken
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	//url := fmt.Sprintf(`https://avito.ru/oauth?response_type=code&client_id=%s`, ClientID)
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//todo struct for data
	json.Marshal(reader)
}
