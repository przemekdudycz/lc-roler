package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"livechat.com/lc-roler/helpers"
	"livechat.com/lc-roler/models"
)

type WebhookPostback struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type WebhookEventBody struct {
	Id       string           `json:"id"`
	Type     string           `json:"type"`
	Message  string           `json:"message"`
	Postback *WebhookPostback `json:"postback"`
}

type EventWebhookRequestPayload struct {
	ChatId string           `json:"chat_id"`
	Event  WebhookEventBody `json:"event"`
}

type EventWebhookRequest struct {
	Payload EventWebhookRequestPayload `json:"payload"`
}

func HandleEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("New event handler \n")

	body, _ := ioutil.ReadAll(r.Body)
	jsonBody := EventWebhookRequest{}
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		fmt.Printf("There was an error: %v \n", err.Error())
	}

	fmt.Printf("Event body: %v", string(body))

	if jsonBody.Payload.Event.Postback.Value == helpers.GetAgentsList {
		handleGetAgentsList(jsonBody.Payload.ChatId)
	}
}

func handleGetAgentsList(chatId string) {
	agentsList := models.GetAgentsList()
	agentListMessage := helpers.AgentsListMessage(chatId, agentsList)
	fmt.Printf("AgentRole: %v", agentsList[0].Role)

	models.SendMessageToCustomer(agentListMessage)
	fmt.Printf("Message sended!")
}
