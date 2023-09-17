package api

import (
	"net/http"
)

func AvitoRouter() {
	http.HandleFunc("/auth", AuthHandler)
	http.HandleFunc("/avito_hook", WebhookHandler)
	http.HandleFunc("/ids", GetListOfResponses)
	http.HandleFunc("/add_notification", SetNotificationEnabled)
	http.HandleFunc("/notifications", GetNotificationsInfo)
	http.HandleFunc("/responses", GetIdsOfResponses)
}
