package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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
