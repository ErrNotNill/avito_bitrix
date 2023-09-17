package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetNotificationsInfo(w http.ResponseWriter, r *http.Request) {
	token := GetToken()
	notification := &Notification{}
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
	json.Unmarshal(newbody, &notification)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newbody)))
	fmt.Println("req.Body", req.Body)
	fmt.Println("token: ", bearer)
	w.Write([]byte(notification.Url))
}

func SetNotificationEnabled(w http.ResponseWriter, r *http.Request) {
	token := GetToken()
	notification := &Notification{}
	r.Form.Add("url", notification.Url)

	fmt.Println("token from DB: ", token)
	urlApi := `https://onviz-api.ru/avito_hook`
	requestBody, err := json.Marshal(map[string]string{"url": urlApi, "secret": "secret"})

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
