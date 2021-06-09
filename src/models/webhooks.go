package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"livechat.com/lc-roler/helpers"
)

type WebhookResponse struct {
	Id            string `json:"id"`
	Url           string `json:"url"`
	Description   string `json:"description"`
	Action        string `json:"action"`
	OwnerClientId string `json:"owner_client_id"`
	Type          string `json:"type"`
}

func GetWebhooksList(clientId string) []WebhookResponse {
	httpClient := GetAuthenticatedHttpClient()
	getWebhooksListUrl := helpers.WebhooksListUrl()

	reqBodyValues := map[string]interface{}{
		"owner_client_id": clientId,
	}
	reqBody, _ := json.Marshal(reqBodyValues)
	req, _ := http.NewRequest("POST", getWebhooksListUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.PerformRequest(req, "agentAuth")

	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	defer res.Body.Close()

	raw, _ := ioutil.ReadAll(res.Body)

	resPayload := []WebhookResponse{}
	jsonerr := json.Unmarshal(raw, &resPayload)
	if jsonerr != nil {
		panic(err)
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

func RegisterWebhook(clientId string, webhookAction string, destinationWebhookUrl string) string {
	httpClient := GetAuthenticatedHttpClient()
	registerWebhookUrl := helpers.RegisterWebhookUrl()

	reqBodyValues := map[string]interface{}{
		"url":             destinationWebhookUrl,
		"description":     "New chat lc-roler webhook",
		"action":          webhookAction,
		"secret_key":      "verysecretkey",
		"owner_client_id": clientId,
		"type":            "license",
	}
	reqBody, _ := json.Marshal(reqBodyValues)
	req, _ := http.NewRequest("POST", registerWebhookUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.PerformRequest(req, "agentAuth")

	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	defer res.Body.Close()
	raw, _ := ioutil.ReadAll(res.Body)
	return string(raw)
}

func EnableWebhooks() {
	httpClient := GetAuthenticatedHttpClient()
	enableWebhookUrl := helpers.EnableWebhookUrl()

	reqBody, _ := json.Marshal(map[string]interface{}{})

	req, _ := http.NewRequest("POST", enableWebhookUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.PerformRequest(req, "agentAuth")
	fmt.Printf("After enabling: %v", res.StatusCode)

	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("Invalid response: %v", res.StatusCode)
	}

	defer res.Body.Close()
}