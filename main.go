package main

import (
	"avito_bitrix/api"
	"fmt"
	"net/http"
)

func main() {

	api.InitDB(`mysqld:mysql@tcp(45.141.79.120:3306)/Onviz`)
	//api.CreateAccessToken() //todo if you need to recreate token

	api.AvitoRouter()
	//handler.SetNotificationEnabled(Token) //todo if you need to change notification url
	fmt.Println("server started")

	http.ListenAndServe(":9090", nil)

}
