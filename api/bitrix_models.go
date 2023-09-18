package api

import "fmt"

type SmartProcess struct {
	EntityTypeId int `json:"entityTypeId"`
	Fields       struct {
		Title string `json:"title"`
		//UfCrm18_1694004396 string `json:"ufCrm18_1694004396"`
	}
}
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
