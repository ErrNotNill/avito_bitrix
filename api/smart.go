package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

/*func SmartProcessCustomFields(fieldId string) {
	customFields := make(map[string]string)
	FindSettingsSettings(fieldId)

}*/

type CustomFields struct {
	EntityTypeId int `json:"entityTypeId"`
	Fields       struct {
		Title   string `json:"title"`
		Address string `json:"ufCrm18_1695020208522"`
		Phone   string `json:"ufCrm18_1695020373419"`
		Name    string `json:"ufCrm18_1695016806010"`
	} `json:"fields"`
}

func SmartProcessFields() string {

	var title string
	var smartId int
	newReq := fmt.Sprintf(`{
"entityTypeId":%v,
"fields":{
"title": "%s",
"ufCrm18_1695020208522": "%s",
"ufCrm18_1695020373419" : "%s",
"ufCrm18_1695016806010" : "%s"
}
}`, smartId, title)
	return newReq
}

func AddSmartProcess(title string, smartId int, address string, phone string, clientName string) {
	newReq := fmt.Sprintf(`{
"entityTypeId":%v,
"fields":{
"title": "%s",
"ufCrm18_1695020208522": "%s",
"ufCrm18_1695020373419" : "%s",
"ufCrm18_1695016806010" : "%s"
}
}`, smartId, title, address, phone, clientName) //UF_CRM_18_1694959668872: address, UF_CRM_18_1694004785: phone, UF_CRM_18_1695016806010: name
	tr := bytes.NewReader([]byte(newReq))
	_, err := http.Post("https://onviz.bitrix24.ru/rest/13938/f72sxd4fph27nuq2/crm.item.add", "application/json", tr)
	if err != nil {
		log.Println("Error http:post request to Bitrix24")
	}
}
