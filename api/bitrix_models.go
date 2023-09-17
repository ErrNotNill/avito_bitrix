package api

type SmartProcess struct {
	EntityTypeId int `json:"entityTypeId"`
	Fields       struct {
		Title string `json:"title"`
		//UfCrm18_1694004396 string `json:"ufCrm18_1694004396"`
	}
}
