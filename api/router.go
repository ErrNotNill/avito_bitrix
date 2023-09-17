package api

import (
	"net/http"
)

func AvitoRouter() {
	http.HandleFunc("/auth", AuthHandler)
	http.HandleFunc("/avito_hook", SetNotificationEnabled) //todo replace with WebhookHandler
	http.HandleFunc("/ids", GetListOfResponses)
	http.HandleFunc("/add_notification", WebhookHandler) //todo replace with SetNotificationEnabled
	http.HandleFunc("/notifications", GetNotificationsInfo)
	http.HandleFunc("/responses", GetIdsOfResponses)
	http.HandleFunc("/get_by_ids", GetByIdsHandler)

}
