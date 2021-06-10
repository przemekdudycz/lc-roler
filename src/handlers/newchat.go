package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"livechat.com/lc-roler/helpers"
	"livechat.com/lc-roler/models"
)

type NewChatWebhookChat struct {
	Id string `json:"id"`
}

type NewChatWebhookRequestPayload struct {
	Chat NewChatWebhookChat `json:"chat"`
}

type NewChatWebhookRequest struct {
	Payload NewChatWebhookRequestPayload `json:"payload"`
}

func HandleNewChat(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("New chat handler \n")
	body, _ := ioutil.ReadAll(r.Body)
	jsonBody := NewChatWebhookRequest{}
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		fmt.Printf("There was an error: %v \n", err.Error())
	}

	chatId := jsonBody.Payload.Chat.Id
	fmt.Printf("ChatId: %v \n", chatId)

	messageBody := helpers.WelcomeMessage(chatId)
	eventId := models.SendMessageToCustomer(messageBody)
	fmt.Printf("EventId: %v \n", eventId)
}
