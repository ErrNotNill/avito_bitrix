package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

type Ids struct {
	Ids string `json:"ids"`
}

func GetIdsOfResponses(w http.ResponseWriter, r *http.Request) {
	token := GetToken()
	fmt.Println("token from DB: ", token)
	var bearer = "Bearer " + token
	url := `https://api.avito.ru/job/v1/applications/get_ids?updatedAtFrom=2023-06-12`
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error")
	}
	req.Header.Add("Authorization", bearer)
	newclient := &http.Client{}
	rez, err := newclient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer rez.Body.Close()
	newbody, err := io.ReadAll(rez.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newbody)))
	fmt.Println("req.Body", req.Body)
	fmt.Println("token: ", bearer)
}

// todo notifications before it
func GetListOfResponses(w http.ResponseWriter, r *http.Request) {
	ids := &Ids{}
	token := GetToken()
	fmt.Println("token from DB: ", token)
	var bearer = "Bearer " + token
	url := `https://api.avito.ru/job/v1/applications/get_by_ids`
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error")
	}
	req.Header.Add("Authorization", bearer)
	newclient := &http.Client{}
	rez, err := newclient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer rez.Body.Close()
	newbody, err := io.ReadAll(rez.Body)
	json.Unmarshal(newbody, &ids)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newbody)))
	fmt.Println("req.Body", req.Body)
	fmt.Println("token: ", bearer)
	fmt.Println("ids.Ids: ", ids.Ids)
}

func GetNotificationsInfo(w http.ResponseWriter, r *http.Request) {
	token := GetToken()
	fmt.Println("token from DB: ", token)
	var bearer = "Bearer " + token
	url := `https://api.avito.ru/job/v1/applications/webhook`
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error")
	}
	req.Header.Add("Authorization", bearer)
	newclient := &http.Client{}
	rez, err := newclient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer rez.Body.Close()
	newbody, err := io.ReadAll(rez.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newbody)))
	fmt.Println("req.Body", req.Body)
	fmt.Println("token: ", bearer)
}

func SetNotificationEnabled(w http.ResponseWriter, r *http.Request) {
	token := GetToken()
	fmt.Println("token from DB: ", token)
	urlApi := `https://onviz-api.ru/add_notification`
	requestBody, err := json.Marshal(map[string]string{"url": urlApi, "secret": "secret"})

	notification := &Notification{}
	var bearer = "Bearer " + token
	url := fmt.Sprintf(`https://api.avito.ru/job/v1/applications/webhook`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	//req.Header.Add("url", urlApi)
	newclient := &http.Client{}
	rez, err := newclient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer rez.Body.Close()
	body, err := io.ReadAll(rez.Body)
	json.Unmarshal(body, &notification)
	fmt.Println("notification.Url", notification.Url)
	fmt.Println("string(body) : ", string(body))
}

func CreateAccessToken() {
	client := &http.Client{}
	auth := &Auth{}
	grantType := "client_credentials"
	url := fmt.Sprintf(`https://api.avito.ru/token?client_id=%s&client_secret=%s&grant_type=%s`, ClientID, ClientSecret, grantType)
	post, err := client.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Println("Failed to create token: ", err)
	}
	body, err := io.ReadAll(post.Body)
	json.Unmarshal(body, &auth)
	fmt.Println("string(body).CreateAccessToken : ", string(body))
	fmt.Println("status code: ", post.StatusCode)
	fmt.Println("auth.AccessToken: ", auth.AccessToken)
	AddToken(auth)
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
	fmt.Println("auth.AccessToken: ", auth.AccessToken)
	//Token = auth.AccessToken
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

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		for {
			reader, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//todo struct for data
			json.Marshal(reader)
			_, err = w.Write(reader)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}
