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

	runAction(jsonBody.Payload.Event.Postback.Value, jsonBody.Payload.ChatId)

	w.Write([]byte{})
}

func runAction(actionValue string, chatId string) {
	if actionValue == helpers.GetAccountsList {
		handleGetAccountsList(chatId)
		return
	}

	splittedActionVal := strings.Split(actionValue, ":")
	actionType := splittedActionVal[0]

	if actionType == helpers.SetRoles {
		splittedVal := strings.Split(actionValue, ":")
		handleSetRoles(chatId, splittedVal[1])
		return
	}

	if actionType == helpers.SetRole {
		splittedVal := strings.Split(actionValue, ":")
		handleSetRole(chatId, splittedVal[1], splittedVal[2], splittedVal[3])
		return
	}

	if actionType == helpers.RevokeRoles {
		splittedVal := strings.Split(actionValue, ":")
		handleRevokeRoles(chatId, splittedVal[1])
	}

	if actionType == helpers.RevokeRole {
		splittedVal := strings.Split(actionValue, ":")
		handleRevokeRole(chatId, splittedVal[1], splittedVal[2], splittedVal[3])
	}
}

func handleGetAccountsList(chatId string) {
	accountsList := models.GetAccountsList()
	accountsListMessage := helpers.AccountsListMessage(chatId, accountsList)

	models.SendMessageToCustomer(accountsListMessage)
	fmt.Printf("Get accounts message sended!")
}

func handleSetRole(chatId string, product string, accountId string, role string) {
	setRoles := []models.Role{
		{Product: product, Role: role},
	}
	models.UpdateAccountRoles(accountId, setRoles, nil)
	fmt.Printf("Roles updated!")

	sendWelcomeMessage(chatId)
}

func handleSetRoles(chatId string, accountId string) {
	account := models.GetAccount(accountId)
	message := helpers.SetAccountRolesMessage(chatId, account)

	models.SendMessageToCustomer(message)
	fmt.Printf("Set roles message sended!")
}

func sendWelcomeMessage(chatId string) {
	welcomeMessage := helpers.WelcomeMessage(chatId)
	models.SendMessageToCustomer(welcomeMessage)
}

func handleRevokeRoles(chatId string, accountId string) {
	account := models.GetAccount(accountId)
	message := helpers.RevokeAccountRolesMessage(chatId, account)

	models.SendMessageToCustomer(message)
}

func handleRevokeRole(chatId string, product string, accountId string, role string) {
	deleteRoles := []models.Role{
		{Product: product, Role: role},
	}
	models.UpdateAccountRoles(accountId, nil, deleteRoles)
	fmt.Printf("Roles updated!")

	sendWelcomeMessage(chatId)
}
