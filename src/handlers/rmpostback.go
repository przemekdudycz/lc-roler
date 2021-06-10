package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookRMPostbackBody struct {
	Id      string `json:"id"`
	Toggled bool   `json:"toggled"`
}

type RMPostbackWebhookRequestPayload struct {
	Postback WebhookRMPostbackBody `json:"postback"`
}

type RMPostbackWebhookRequest struct {
	ChatId  string                          `json:"chat_id"`
	Payload RMPostbackWebhookRequestPayload `json:"payload"`
}

func HandleRMPostback(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("RichMessage postback handler \n")

	body, _ := ioutil.ReadAll(r.Body)
	jsonBody := RMPostbackWebhookRequest{}
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		fmt.Printf("There was an error: %v \n", err.Error())
	}

	fmt.Printf("Postback body: %v", string(body))
}
