package main

import (
	"avito_bitrix/avito"
	"avito_bitrix/avito/handler"
	"fmt"
	"net/http"
)

func main() {

	return

	handler.FindSubstr("ClientID")

	return

	handler.InitDB(`mysqld:mysql@tcp(45.141.79.120:3306)/Onviz`)
	//handler.CreateAccessToken()
	//urlDb := os.Getenv("URL_MYSQL")
	//bitrixDomain := os.Getenv("URL_MYSQL")
	avito.AvitoRouter()
	//handler.SetNotificationEnabled(Token)
	fmt.Println("server started")
	http.ListenAndServe(":9090", nil)
}
