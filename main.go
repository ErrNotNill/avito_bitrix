package main

import (
	"avito_bitrix/api"
	"fmt"
	"net/http"
)

func main() {

	/*jsonStr := `{"applies":[{"id":"650749d1e3ab7b1a5f1f46df","negotiation_id":501839919,"type":"by_chat","created_at":"2023-09-17T18:47:44Z","updated_at":"2023-09-17T18:47:50Z","prevalidation":{"status":"enrichment_finished","summary":[{"label":"Пол","value":"Мужчина","variable":"job_gender"},{"label":"Телефон","value":"+79991211232","variable":"job_phone"},{"label":"Возраст","value":"43 года","variable":"widget_id_496964"},{"label":"ФИО","value":"Михалов Валерий Валерьевич","variable":"job_fio"},{"label":"Гражданство","value":"Россия","variable":"widget_id_496966"},{"label":"Опыт работы по профессии","value":"нет","variable":"widget_id_886567"}]},"applicant":{"id":"fdb4ce70-ef19-4b9e-a222-8a9b91a5ebd6","data":{"name":"Услуги"}},"contacts":{"chat":{"value":"u2i-X8rfCxVxOjF8J4Skj4W0pA"},"phones":[{"value":"79536852874","status":null}]},"vacancy_id":3035894401,"employee_id":null}]}`

	var root api.Root
	if err := json.Unmarshal([]byte(jsonStr), &root); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Access the parsed data
	for _, apply := range root.Applies {
		fmt.Println("ID:", apply.ID)
		fmt.Println("Negotiation ID:", apply.NegotiationID)
		fmt.Println("Type:", apply.Type)
		fmt.Println("Created At:", apply.CreatedAt)
		fmt.Println("Updated At:", apply.UpdatedAt)
		fmt.Println("Prevalidation Status:", apply.Prevalidation.Status)
		fmt.Println("Applicant ID:", apply.Applicant.ID)
		fmt.Println("Applicant Name:", apply.Applicant.Data.Name)
		fmt.Println("Chat Value:", apply.Contacts.Chat.Value)
		fmt.Println("Phone Value:", apply.Contacts.Phones[0].Value)
		fmt.Println("Vacancy ID:", apply.VacancyID)
		fmt.Println("Employee ID:", apply.EmployeeID)
		fmt.Println()
	}*/

	api.InitDB(`mysqld:mysql@tcp(45.141.79.120:3306)/Onviz`)
	//handler.CreateAccessToken()
	//urlDb := os.Getenv("URL_MYSQL")
	//bitrixDomain := os.Getenv("URL_MYSQL")
	api.AvitoRouter()
	//handler.SetNotificationEnabled(Token)
	fmt.Println("server started")
	http.ListenAndServe(":9090", nil)

}
