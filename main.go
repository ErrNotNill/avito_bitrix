package main

import (
	"avito_bitrix/api"
	"fmt"
	"net/http"
)

func main() {

	api.InitDB(`mysqld:mysql@tcp(45.141.79.120:3306)/Onviz`)
	//handler.CreateAccessToken()
	//urlDb := os.Getenv("URL_MYSQL")
	//bitrixDomain := os.Getenv("URL_MYSQL")
	api.AvitoRouter()
	//handler.SetNotificationEnabled(Token)
	fmt.Println("server started")
	http.ListenAndServe(":9090", nil)

}
