package api

import "time"

type PrevalidationSummaryApplies struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Variable string `json:"variable"`
}

type PrevalidationApplies struct {
	Status  string                 `json:"status"`
	Summary []PrevalidationSummary `json:"summary"`
}

type ApplicantData struct {
	Name string `json:"name"`
}

type ApplicantApplies struct {
	ID   string        `json:"id"`
	Data ApplicantData `json:"data"`
}

type Chat struct {
	Value string `json:"value"`
}

type Phone struct {
	Value  string `json:"value"`
	Status string `json:"status"`
}

type ContactsApplies struct {
	Chat   Chat    `json:"chat"`
	Phones []Phone `json:"phones"`
}

type Apply struct {
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

type RootApplies struct {
	Applies []Apply `json:"applies"`
}
