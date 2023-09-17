package main

import (
	"avito_bitrix/avito"
	"avito_bitrix/avito/handler"
	"net/http"
)

func main() {

	handler.InitDB(`mysqld:mysql@tcp(45.141.79.120:3306)/Onviz`)
	//handler.CreateAccessToken()
	//urlDb := os.Getenv("URL_MYSQL")
	//bitrixDomain := os.Getenv("URL_MYSQL")

	avito.AvitoRouter()
	//handler.SetNotificationEnabled(Token)

	http.ListenAndServe(":8080", nil)

}
