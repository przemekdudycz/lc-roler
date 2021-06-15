package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	Text     string           `json:"text"`
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

	if jsonBody.Payload.Event.Text == helpers.AppStartUp {
		sendWelcomeMessage(jsonBody.Payload.ChatId)
		w.Write([]byte{})
		return
	}

	if jsonBody.Payload.Event.Postback.Value == helpers.GetAgentsList {
		handleGetAgentsList(jsonBody.Payload.ChatId)
		w.Write([]byte{})
		return
	}

	if strings.Contains(jsonBody.Payload.Event.Postback.Value, helpers.UseViceOwnerRole) ||
		strings.Contains(jsonBody.Payload.Event.Postback.Value, helpers.UseAdministratorRole) ||
		strings.Contains(jsonBody.Payload.Event.Postback.Value, helpers.UseNormalRole) {
		splittedValue := strings.Split(jsonBody.Payload.Event.Postback.Value, ":")
		handleRoleChange(splittedValue[0], splittedValue[1])
		sendWelcomeMessage(jsonBody.Payload.ChatId)
		w.Write([]byte{})
		return
	}
}

func handleGetAgentsList(chatId string) {
	agentsList := models.GetAgentsList()
	agentListMessage := helpers.AgentsListMessage(chatId, agentsList)
	fmt.Printf("AgentRole: %v", agentsList[0].Role)

	models.SendMessageToCustomer(agentListMessage)
	fmt.Printf("Message sended!")
}

func handleRoleChange(agentId string, roleName string) {
	models.UpdateAgentRole(roleName, agentId)
	fmt.Println("Role is updated!")
}

func sendWelcomeMessage(chatId string) {
	welcomeMessage := helpers.WelcomeMessage(chatId)
	models.SendMessageToCustomer(welcomeMessage)
}
