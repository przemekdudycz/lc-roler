package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	res, err := httpClient.PerformRequest(req, "agentAuth")
	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("Send message error: %v \n", res.StatusCode)
	}

	defer res.Body.Close()
	raw, _ := ioutil.ReadAll(res.Body)

	return string(raw)
}
