package api

type Ids struct {
	Ids string `json:"ids"`
}

type Vacancy struct {
	Title  string `json:"title"`
	Params struct {
		Address string `json:"address"`
	} `json:"params"`
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

type Response struct {
	ApplyId   string `json:"applyId"`
	VacancyId string `json:"vacancy_id"`
}

var (
	ApplyId   string
	VacancyId string
)
