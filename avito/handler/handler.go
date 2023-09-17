package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func FindSubstr(substr string) string {
	var substring string
	file, err := os.ReadFile("custom_fields")
	if err != nil {
		fmt.Println("Error reading")
	}
	list := strings.Split(string(file), ",")
	for k, _ := range list {
		withoutSpaces := strings.Join(strings.Fields(list[k]), "")
		if strings.Contains(withoutSpaces, substr) {
			pattern := `<([^>]+)>`
			re := regexp.MustCompile(pattern)
			match := re.FindString(withoutSpaces)
			if match != "" {
				// Remove the "<" and ">" symbols
				substring = match[1 : len(match)-1]
				fmt.Println("Substring:", substring)
			} else {
				fmt.Println("No match found")
			}
		}

	}
	return substring
}

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

type Vacancy struct {
	Title  string `json:"title"`
	Params struct {
		Address string `json:"address"`
	} `json:"params"`
}

func GetVacancyInfo(vacancyId string) *Vacancy {
	token := GetToken()
	vacancy := &Vacancy{}
	fmt.Println("token from DB: ", token)
	var bearer = "Bearer " + token
	url := `https://api.avito.ru/job/v2/vacancies/` + vacancyId
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
	json.Unmarshal(newbody, &vacancy)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newbody)))
	fmt.Println("vacancy.Params.Address", vacancy.Params.Address)
	fmt.Println("vacancy.Title: ", vacancy.Title)
	return vacancy
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
	urlApi := `https://onviz-api.ru/add_notification`
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

type Response struct {
	ApplyId   string `json:"applyId"`
	VacancyId string `json:"vacancy_id"`
}

var (
	ApplyId   string
	VacancyId string
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	response := &Response{}
	if r.Method == "POST" {
		for {
			reader, err := io.ReadAll(r.Body)
			json.Unmarshal(reader, &response)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ApplyId = response.ApplyId
			VacancyId = response.VacancyId
			vacancy := GetVacancyInfo(VacancyId)
			AddSmartProcess(vacancy.Title, 139, vacancy.Params.Address)
		}
	}
}

type User struct {
	Data struct {
		Name string `json:"name"`
	} `json:"data"`
	Contacts struct {
		Chat struct {
			Value string `json:"value"`
		} `json:"chat"`
	} `json:"contacts"`
	Phones struct {
		Value  string `json:"value"`
		Status string `json:"status"`
	}
}

type SmartProcess struct {
	EntityTypeId int `json:"entityTypeId"`
	Fields       struct {
		Title string `json:"title"`
		//UfCrm18_1694004396 string `json:"ufCrm18_1694004396"`
	}
}

/*
 "item": {
    "id": 12,
    "xmlId": null,
    "title": "TEST",
    "createdBy": 13938,
    "updatedBy": 13938,
    "movedBy": 13938,
    "createdTime": "2023-09-17T16:11:00+03:00",
    "updatedTime": "2023-09-17T16:11:00+03:00",
    "movedTime": "2023-09-17T16:11:00+03:00",
    "categoryId": 26,
    "opened": "N",
    "stageId": "DT139_26:NEW",
    "previousStageId": "",
    "begindate": "2023-09-17T03:00:00+03:00",
    "closedate": "2023-09-24T03:00:00+03:00",
    "companyId": null,
    "contactId": null,
    "opportunity": 0,
    "isManualOpportunity": "N",
    "taxValue": 0,
    "currencyId": "RUB",
    "mycompanyId": 14,
    "sourceId": "OTHER",
    "sourceDescription": null,
    "webformId": null,
    "ufCrm18_1694004396": "",
    "ufCrm18_1694004544": "",
    "ufCrm18_1694004572": "",
    "ufCrm18_1694004582": "",
    "ufCrm18_1694004588": "",
    "ufCrm18_1694004595": "",
    "ufCrm18_1694004606": "",
    "ufCrm18_1694004655": "",
    "ufCrm18_1694004761": "",
    "ufCrm18_1694004785": "",
    "ufCrm18_1694004814": "",
    "ufCrm18_1694004929": "",
    "ufCrm18_1694005068": "",
    "ufCrm18_1694005075": "",
    "ufCrm18_1694005082": "",
    "ufCrm18_1694005088": "",
    "assignedById": 13938,
    "lastActivityBy": 13938,
    "lastActivityTime": "2023-09-17T16:11:00+03:00",
    "parentId2": null,
    "utmSource": null,
    "utmMedium": null,
    "utmCampaign": null,
    "utmContent": null,
    "utmTerm": null,
    "observers": [],
    "contactIds": [],
    "entityTypeId": 139
*/

/*func SmartProcessCustomFields(fieldId string) {
	customFields := make(map[string]string)
	FindSubstr(fieldId)

}*/

func SmartProcessFields() string {
	var title string
	var smartId int
	newReq := fmt.Sprintf(`{
entityTypeId:%s,
fields:{
title: %s,
}
}`, smartId, title)
	return newReq
}

func AddSmartProcess(title string, smartId int, address string) {
	newReq := fmt.Sprintf(`{
entityTypeId:%s,
fields:{
title: %s,
UF_CRM_18_1694959668872: %s
}
}`, smartId, title, address)
	tr := bytes.NewReader([]byte(newReq))
	_, err := http.Post("https://onviz.bitrix24.ru/rest/13938/f72sxd4fph27nuq2/crm.item.add", "application/json", tr)
	if err != nil {
		log.Println("Error http:post request to Bitrix24")
	}
}
