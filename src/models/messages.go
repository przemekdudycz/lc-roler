package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"livechat.com/lc-roler/config"
)

type ChatEvent struct {
	Type       string                   `json:"type"`
	TemplateId string                   `json:"template_id"`
	Elements   []map[string]interface{} `json:"elements"`
}

type ChatMessage struct {
	ChatId string    `json:"chat_id"`
	Event  ChatEvent `json:"event"`
}

func SendMessageToCustomer(reqBodyValues ChatMessage) string {
	httpClient := GetAuthenticatedHttpClient()
	authConfig := config.GetAuthConfiguration()
	sendEventUrl := config.GetSendEventUrl(authConfig.LicenseId)

	reqBody, _ := json.Marshal(reqBodyValues)
	req, _ := http.NewRequest("POST", sendEventUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	var eventIdResponse = struct {
		EventId string `json:"event_id"`
	}{}
	err := httpClient.SendRequest(req, "agentAuth", &eventIdResponse)

	if err != nil {
		fmt.Printf("there was an error when SendMessageToCustomer: %v", err.Error())
	}

	return eventIdResponse.EventId
}
