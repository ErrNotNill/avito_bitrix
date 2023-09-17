package avito

import (
	"avito_bitrix/avito/handler"
	"net/http"
)

func AvitoRouter() {
	http.HandleFunc("/auth", handler.AuthHandler)
	http.HandleFunc("/avito_hook", handler.WebhookHandler)
	http.HandleFunc("/ids", handler.GetListOfResponses)
	http.HandleFunc("/add_notification", handler.SetNotificationEnabled)
	http.HandleFunc("/notifications", handler.GetNotificationsInfo)
	http.HandleFunc("/responses", handler.GetIdsOfResponses)
}
