package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FindCustomFields(substr string) string {
	var substring string
	file, err := os.ReadFile("crm_fields")
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
				//fmt.Println("Substring:", substring)
			} else {
				fmt.Println("No match found")
			}
		}
	}
	return substring
}

func FindSettings(substr string) string {
	var substring string
	file, err := os.ReadFile("settings")
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
				//fmt.Println("Substring:", substring)
			} else {
				fmt.Println("No match found")
			}
		}
	}
	return substring
}

func GetVacancyInfo(vacancyId int) *Vacancy {
	//token := GetToken()
	newToken := CreateAccessToken()

	vacancy := &Vacancy{}
	//fmt.Println("token from DB: ", token)
	var bearer = "Bearer " + newToken
	url := `https://api.avito.ru/job/v2/vacancies/` + strconv.Itoa(vacancyId)
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
	//log.Println("newBody", string([]byte(newbody)))
	//fmt.Println("vacancy.Params.Address", vacancy.Params.Address)
	//fmt.Println("vacancy.Title: ", vacancy.Title)
	return vacancy
}

func GetIdsOfResponses(w http.ResponseWriter, r *http.Request) {
	//token := GetToken()
	newToken := CreateAccessToken()

	//fmt.Println("token from DB: ", token)
	w.Write([]byte("use query params: {date:2023-06-12}"))

	date := r.FormValue("date")
	var bearer = "Bearer " + newToken
	url := `https://api.avito.ru/job/v1/applications/get_ids?updatedAtFrom=` + date
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
	//newbody, err := io.ReadAll(rez.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	//log.Println("newBody", string([]byte(newbody)))
	//fmt.Println("req.Body", req.Body)
	//fmt.Println("token: ", bearer)
}

/*func GetInfoAboutAdvertisement() {
	ids := &Ids{}
	newReq := fmt.Sprintf(`{"ids": "%s"}`, applyId)
	tr := bytes.NewReader([]byte(newReq))
	token := GetToken()
	fmt.Println("token from DB: ", token)
	var bearer = "Bearer " + token
	url := `https://api.avito.ru/job/v1/applications/get_by_ids`
	req, err := http.NewRequest("POST", url, tr)
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
}*/

func GetByIdsHandler(w http.ResponseWriter, r *http.Request) {
	ids := &Ids{}
	applyId := r.FormValue("apply_id") //"650721b4e3ab7b1a5fe07c85"
	newReq := fmt.Sprintf(`{"ids": ["%s"]}`, applyId)
	tr := bytes.NewReader([]byte(newReq))
	//token := GetToken()
	newToken := CreateAccessToken()

	fmt.Println("token from DB: ", newToken)
	var bearer = "Bearer " + newToken
	url := `https://api.avito.ru/job/v1/applications/get_by_ids`
	req, err := http.NewRequest("POST", url, tr)
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

func GetByIds(applyId string) (int, string, string, string, string) {
	var root Root
	//vacancyResp := &VacancyResponse{}
	newReq := fmt.Sprintf(`{"ids": ["%s"]}`, applyId)
	tr := bytes.NewReader([]byte(newReq))
	//token := GetToken()
	newToken := CreateAccessToken()

	//fmt.Println("token from DB: ", token)
	var bearer = "Bearer " + newToken
	url := `https://api.avito.ru/job/v1/applications/get_by_ids`
	req, err := http.NewRequest("POST", url, tr)
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

	if err := json.Unmarshal(newbody, &root); err != nil {
		fmt.Println("Error while reading the response bytes:", err)
	}
	for _, apply := range root.Applies {
		//	fmt.Println("VacancyID:", apply.VacancyID)
		//	fmt.Println("apply.Applicant.Data.Name:", apply.Applicant.Data.Name)
		//	fmt.Println("apply.Contacts.Phones:", apply.Contacts.Phones)
		var phoneValue string
		for _, phone := range apply.Contacts.Phones {
			phoneValue = phone.Value
		}

		return apply.VacancyID, apply.Applicant.Data.Name, phoneValue, apply.Applicant.ID, apply.Contacts.Chat.Value
	}
	//	log.Println("newBody from GetByIds: ", string([]byte(newbody)))
	err = os.WriteFile("response", []byte(newbody), os.FileMode(0644))
	if err != nil {
		fmt.Println("Error while writing the response")
	}
	readFile, err := os.ReadFile("response")
	fmt.Println(string(readFile))
	//fmt.Println("req.Body GetByIds", req.Body)
	return 0, "", "", "", ""
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Secret") == "secret" {
		response := &Response{}
		if r.Method == "POST" {
			reader, err := io.ReadAll(r.Body)
			log.Println("newBody WebhookHandler: ", string([]byte(reader)))

			err = json.Unmarshal(reader, &response)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ApplyId = response.ApplyId
			vacancyId, Name, Phone, applies, chat := GetByIds(ApplyId)
			fmt.Println("id_applies", applies)
			fmt.Println("id_vacancy", vacancyId)
			fmt.Println("Имя: ", Name)
			fmt.Println("Телефон: ", Phone)
			fmt.Println("chat_id: ", chat)

			vac := GetVacancyInfo(vacancyId)
			fmt.Println("Название вакансии: ", vac.Title)
			fmt.Println("Адрес вакансии: ", vac.Params.Address)
			AddSmartProcess(vac.Title, 139, vac.Params.Address, Phone, Name, applies, chat)
		}
	}
}
