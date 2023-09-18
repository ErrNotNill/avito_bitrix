package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func AddSmartProcess(title string, smartId int, address string, phone string, clientName string) {
	phoneField := FindCustomFields("Телефон")
	addressField := FindCustomFields("Адрес")
	clientNameField := FindCustomFields("Имя")
	newReq := fmt.Sprintf(`{
"entityTypeId":%v,
"fields":{
"title": "%s",
"%s": "%s",
"%s" : "%s",
"%s" : "%s"
}
}`, smartId, title, addressField, address, phoneField, phone, clientNameField, clientName) //UF_CRM_18_1694959668872: address, UF_CRM_18_1694004785: phone, UF_CRM_18_1695016806010: name
	tr := bytes.NewReader([]byte(newReq))
	_, err := http.Post("https://onviz.bitrix24.ru/rest/13938/f72sxd4fph27nuq2/crm.item.add", "application/json", tr)
	if err != nil {
		log.Println("Error http:post request to Bitrix24")
	}
}
