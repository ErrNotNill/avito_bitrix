package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

/*func SmartProcessCustomFields(fieldId string) {
	customFields := make(map[string]string)
	FindSubstr(fieldId)

}*/

func SmartProcessFields() string {
	var title string
	var smartId int
	newReq := fmt.Sprintf(`{entityTypeId:%s,fields:{
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
