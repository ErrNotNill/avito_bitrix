package api

import (
	"time"
)

type Ids struct {
	Ids string `json:"ids"`
}

type Vacancy struct {
	Title  string `json:"title"`
	Params struct {
		Address string `json:"address"`
	} `json:"params"`
}

type VacancyResponse struct {
	Licant struct {
		ID   string `json:"id"`
		Data struct {
			Name string `json:"name"`
		} `json:"data"`
		Contacts struct {
			Chat struct {
				Value  string `json:"value"`
				Phones []struct {
					Value  string `json:"value"`
					Status string `json:"status"`
				} `json:"phones"`
			}
		}
		VacancyID   int64 `json:"vacancy_id"`
		Employee_id int64 `json:"employee_id"`
	} `json:"licant"`
	/*licant":{"id":"fdb4ce70-ef19-4b9e-a222-8a9b91a5ebd6","data":{"name":"Услуги"}},"contacts":{"chat":{"value":"u2i-Ivgfe~_EgbEL2uLzXfThGw"},"phones":[{"value":"79536852874
	","status":null}]},"vacancy_id":3355908978,"employee_id":null}]}*/

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

// new
type PrevalidationSummary struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Variable string `json:"variable"`
}

type Prevalidation struct {
	Status  string                 `json:"status"`
	Summary []PrevalidationSummary `json:"summary"`
}

type Applicant struct {
	ID   string `json:"id"`
	Data struct {
		Name string `json:"name"`
	} `json:"data"`
}

type Contacts struct {
	Chat struct {
		Value string `json:"value"`
	} `json:"chat"`
	Phones []struct {
		Value  string `json:"value"`
		Status string `json:"status"`
	} `json:"phones"`
}

type Applies struct {
	ID            string        `json:"id"`
	NegotiationID int           `json:"negotiation_id"`
	Type          string        `json:"type"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Prevalidation Prevalidation `json:"prevalidation"`
	Applicant     Applicant     `json:"applicant"`
	Contacts      Contacts      `json:"contacts"`
	VacancyID     int           `json:"vacancy_id"`
	EmployeeID    interface{}   `json:"employee_id"`
}

type Root struct {
	Applies []Applies `json:"applies"`
}
