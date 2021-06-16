package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"livechat.com/lc-roler/config"
)

type WebhookResponse struct {
	Id            string `json:"id"`
	Url           string `json:"url"`
	Description   string `json:"description"`
	Action        string `json:"action"`
	OwnerClientId string `json:"owner_client_id"`
	Type          string `json:"type"`
}

type WebhookFilter struct {
	AuthorType string `json:"author_type"`
}

type WebhookRequestBody struct {
	Url           string        `json:"url"`
	Description   string        `json:"description"`
	Action        string        `json:"action"`
	SecretKey     string        `json:"secret_key"`
	OwnerClientId string        `json:"owner_client_id"`
	Type          string        `json:"type"`
	Filters       WebhookFilter `json:"filters"`
}

func GetWebhooksList(clientId string) []WebhookResponse {
	httpClient := GetAuthenticatedHttpClient()
	getWebhooksListUrl := config.WebhooksListUrl()

	reqBodyValues := map[string]interface{}{
		"owner_client_id": clientId,
	}
	reqBody, _ := json.Marshal(reqBodyValues)
	req, _ := http.NewRequest("POST", getWebhooksListUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	resPayload := []WebhookResponse{}
	err := httpClient.SendRequest(req, "agentAuth", &resPayload)

	if err != nil {
		fmt.Printf("there was an error when GetWebhookList: %v", err.Error())
	}

	return resPayload
}

func IsWebhookWithActionInSlice(webhooks []WebhookResponse, webhookAction string) bool {
	for _, webhook := range webhooks {
		if webhook.Action == webhookAction {
			return true
		}
	}
	return false
}

func RegisterWebhook(webhookData WebhookRequestBody) string {
	httpClient := GetAuthenticatedHttpClient()
	registerWebhookUrl := config.RegisterWebhookUrl()
	reqBody, _ := json.Marshal(webhookData)
	req, _ := http.NewRequest("POST", registerWebhookUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	idResponse := struct {
		Id string `json:"id"`
	}{}
	err := httpClient.SendRequest(req, "agentAuth", &idResponse)

	if err != nil {
		fmt.Printf("there was an error when RegisterWebhook: %v", err.Error())
	}

	return idResponse.Id
}

func EnableWebhooks() {
	httpClient := GetAuthenticatedHttpClient()
	enableWebhookUrl := config.EnableWebhookUrl()

	reqBody, _ := json.Marshal(map[string]interface{}{})

	req, _ := http.NewRequest("POST", enableWebhookUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	err := httpClient.SendRequest(req, "agentAuth", nil)

	if err != nil {
		fmt.Printf("there was an error when EnableWebhooks: %v", err.Error())
	}
}
