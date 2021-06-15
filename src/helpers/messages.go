package helpers

import (
	"fmt"

	"livechat.com/lc-roler/models"
)

var products []string = []string{"Accounts", "LiveChat", "HelpDesk"}
var productsRoles map[string][]string = map[string][]string{
	"Accounts": {"administrator", "member", "billing_editor"},
	"LiveChat": {"administrator", "normal"},
	"HelpDesk": {"administrator", "normal", "viewer"},
}

func WelcomeMessage(chatId string) models.ChatMessage {
	return models.ChatMessage{
		ChatId: chatId,
		Event: models.ChatEvent{
			Type:       "rich_message",
			TemplateId: "quick_replies",
			Elements: []map[string]interface{}{
				{
					"title": "What do you want from lc-roler?",
					"buttons": []map[string]interface{}{
						{
							"type":        "message",
							"text":        "Accounts List",
							"postback_id": "send_message",
							"value":       GetAccountsList,
							"user_ids":    []interface{}{},
						},
						{
							"type":        "message",
							"text":        "Nothing",
							"postback_id": "send_message",
							"value":       "Nothing",
							"user_ids":    []interface{}{},
						},
					},
				},
			},
		},
	}
}

func AccountsListMessage(chatId string, accountsList []models.Account) models.ChatMessage {
	message := models.ChatMessage{
		ChatId: chatId,
		Event: models.ChatEvent{
			Type:       "rich_message",
			TemplateId: "cards",
			Elements:   []map[string]interface{}{},
		},
	}

	for _, account := range accountsList {
		message.Event.Elements = append(message.Event.Elements, map[string]interface{}{
			"title":    fmt.Sprintf("Account: %v", account.Name),
			"subtitle": fmt.Sprintf("Id: %v \n Email: %v \n", account.AccountId, account.Email),
			"buttons": []map[string]interface{}{
				{
					"type":        "message",
					"text":        "Set roles",
					"postback_id": "send_message",
					"value":       fmt.Sprintf("%v:%v", SetRoles, account.AccountId),
					"user_ids":    []interface{}{},
				},
			},
		})
	}
	return message
}

func AccountRolesMessage(chatId string, account models.Account) models.ChatMessage {
	message := models.ChatMessage{
		ChatId: chatId,
		Event: models.ChatEvent{
			Type:       "rich_message",
			TemplateId: "cards",
			Elements:   []map[string]interface{}{},
		},
	}

	for _, product := range products {
		if productRoles, productRolesString := filterProductRoles(account.Roles, product); len(productRoles) > 0 {
			message.Event.Elements = append(message.Event.Elements, map[string]interface{}{
				"title":    fmt.Sprintf("Product: %v", product),
				"subtitle": fmt.Sprintf("Roles: \n %v", productRolesString),
				"buttons":  buildRolesButtonsForProduct(product, account),
			})
		}
	}

	return message
}

func filterProductRoles(roles []models.Role, product string) ([]models.Role, string) {
	productRoles := []models.Role{}
	productRolesString := ""
	for _, role := range roles {
		if role.Product == product {
			productRoles = append(productRoles, role)
			productRolesString += fmt.Sprintf("%v, ", role.Role)
		}
	}
	return productRoles, productRolesString
}

func buildRolesButtonsForProduct(product string, account models.Account) []map[string]interface{} {
	buttons := []map[string]interface{}{}
	for _, role := range productsRoles[product] {
		buttons = append(buttons, map[string]interface{}{
			"type":        "message",
			"text":        role,
			"postback_id": "send_message",
			"value":       fmt.Sprintf("%v:%v:%v:%v", SetRole, product, account.AccountId, role),
			"user_ids":    []interface{}{},
		})
	}
	return buttons
}
