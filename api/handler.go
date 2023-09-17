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
	"strings"
	"time"
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
	applyId := "650721b4e3ab7b1a5fe07c85"
	newReq := fmt.Sprintf(`{"ids": ["%s"]}`, applyId)
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
}

type Applicant struct {
	Applicant []struct {
		ID string `json:"id"`
	} `json:"applies"`
}

func GetByIds(applyId string) {
	//vacancyResp := &VacancyResponse{}
	applies := &Applies{}
	newReq := fmt.Sprintf(`{"ids": ["%s"]}`, applyId)
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
	/*licant":{"id":"fdb4ce70-ef19-4b9e-a222-8a9b91a5ebd6","data":{"name":"Услуги"}},"contacts":{"chat":{"value":"u2i-Ivgfe~_EgbEL2uLzXfThGw"},"phones":[{"value":"79536852874
	","status":null}]},"vacancy_id":3355908978,"employee_id":null}]}*/
	json.Unmarshal(newbody, &applies)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody from GetByIds: ", string([]byte(newbody)))
	err = os.WriteFile("response", []byte(newbody), os.FileMode(0644))
	if err != nil {
		fmt.Println("Error while writing the response")
	}
	readFile, err := os.ReadFile("response")
	fmt.Println(string(readFile))

	fmt.Println("req.Body GetByIds", req.Body)
	fmt.Println("applies.VacancyId: ", applies.VacancyID)

}

type ApplicantNew struct {
	Id            string `json:"id"`
	NegotiationId int    `json:"negotiation_id"`
	Type          string `json:"type"`
	Applicant     struct {
		Id       string `json:"id"`
		ResumeId string `json:"resume_id"`
		Data     struct {
			Name        string `json:"name"`
			Gender      string `json:"gender"`
			Citizenship string `json:"citizenship"`
		} `json:"data"`
	} `json:"applicant"`
	VacancyId int64 `json:"vacancy_id"`
}

type Applies struct {
	VacancyID int64 `json:"vacancy_id"`
	Applies   []struct {
		ID            string    `json:"id"`
		NegotiationID int       `json:"negotiation_id"`
		Type          string    `json:"type"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Prevalidation struct {
			Status  string `json:"status"`
			Summary []struct {
				Label    string `json:"label"`
				Value    string `json:"value"`
				Variable string `json:"variable"`
			} `json:"summary"`
		} `json:"prevalidation"`
		Applicant struct {
			ID       string `json:"id"`
			ResumeID string `json:"resume_id"`
			Data     struct {
				Name        string `json:"name"`
				Gender      string `json:"gender"`
				Citizenship string `json:"citizenship"`
			} `json:"data"`
		} `json:"applicant"`
		Contacts struct {
			Chat struct {
				Value string `json:"value"`
			} `json:"chat"`
			Phones []struct {
				Value  string `json:"value"`
				Status any    `json:"status"`
			} `json:"phones"`
		} `json:"contacts"`
		VacancyID  int64 `json:"vacancy_id"`
		EmployeeID any   `json:"employee_id"`
	} `json:"applies"`
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
			fmt.Println("ApplyId", ApplyId)
			GetByIds(ApplyId)

			vacancy := GetVacancyInfo(response.VacancyId)
			AddSmartProcess(vacancy.Title, 139, vacancy.Params.Address)
		}
	}
}
