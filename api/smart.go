package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func AddSmartProcess(title string, smartId int, address string, phone string, clientName string, appliesId string, chatId string) {
	phoneField := FindCustomFields("Телефон")
	addressField := FindCustomFields("Адрес")
	clientNameField := FindCustomFields("Имя")
	appliesIdField := FindCustomFields("Откликнувшийся")
	chatIdField := FindCustomFields("Чат")
	titleVacancyField := FindCustomFields("Название")
	newReq := fmt.Sprintf(`{
"entityTypeId":%v,
"fields":{
"title": "%s",
"%s": "%s",
"%s" : "%s",
"%s" : "%s",
"%s" : "%s",
"%s" : "%s",
"%s" : "%s"
}
}`, smartId, title, addressField, address, phoneField, phone, clientNameField, clientName, appliesIdField, appliesId, chatIdField, chatId, titleVacancyField, title) //UF_CRM_18_1694959668872: address, UF_CRM_18_1694004785: phone, UF_CRM_18_1695016806010: name
	tr := bytes.NewReader([]byte(newReq))

	webhookUrl := FindSettings("WebHookBitrixUrl")
	endUrl := webhookUrl + "crm.item.add"
	_, err := http.Post(endUrl, "application/json", tr)
	if err != nil {
		log.Println("Error http:post request to Bitrix24")
	}
}
