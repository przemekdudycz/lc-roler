package helpers

import (
	"fmt"
	"strings"

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
							"text":        "Manage roles",
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
					"text":        "Assign roles",
					"postback_id": "send_message",
					"value":       fmt.Sprintf("%v:%v", SetRoles, account.AccountId),
					"user_ids":    []interface{}{},
				},
				{
					"type":        "message",
					"text":        "Revoke roles",
					"postback_id": "send_message",
					"value":       fmt.Sprintf("%v:%v", RevokeRoles, account.AccountId),
					"user_ids":    []interface{}{},
				},
			},
		})
	}
	return message
}

func SetAccountRolesMessage(chatId string, account models.Account) models.ChatMessage {
	message := models.ChatMessage{
		ChatId: chatId,
		Event: models.ChatEvent{
			Type:       "rich_message",
			TemplateId: "cards",
			Elements:   []map[string]interface{}{},
		},
	}

	for _, product := range products {
		_, productRolesString := filterProductRoles(account.Roles, product)

		var card map[string]interface{}
		if product == "Accounts" && strings.Contains(productRolesString, "owner") {
			card = map[string]interface{}{
				"title":    fmt.Sprintf("Product: %v", product),
				"subtitle": fmt.Sprintf("Roles: \n %v", productRolesString),
			}
		} else {
			card = map[string]interface{}{
				"title":    fmt.Sprintf("Product: %v", product),
				"subtitle": fmt.Sprintf("Roles: \n %v", productRolesString),
				"buttons":  buildRolesButtonsForProduct(product, account),
			}
		}

		message.Event.Elements = append(message.Event.Elements, card)
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

	if productRolesString == "" {
		productRolesString = "No role assigned"
	} else {
		productRolesString = productRolesString[:len(productRolesString)-2]
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

func RevokeAccountRolesMessage(chatId string, account models.Account) models.ChatMessage {
	message := models.ChatMessage{
		ChatId: chatId,
		Event: models.ChatEvent{
			Type:       "rich_message",
			TemplateId: "cards",
			Elements:   []map[string]interface{}{},
		},
	}

	for _, product := range products {
		productRoles, productRolesString := filterProductRoles(account.Roles, product)

		var card map[string]interface{}
		if len(productRoles) == 0 || (product == "Accounts" && strings.Contains(productRolesString, "owner")) {
			card = map[string]interface{}{
				"title":    fmt.Sprintf("Product: %v", product),
				"subtitle": fmt.Sprintf("Roles: \n %v", productRolesString),
			}
		} else {
			card = map[string]interface{}{
				"title":    fmt.Sprintf("Product: %v", product),
				"subtitle": fmt.Sprintf("Roles: \n %v", productRolesString),
				"buttons":  buildRevokeRolesButtons(product, productRoles, account),
			}
		}

		message.Event.Elements = append(message.Event.Elements, card)
	}

	return message
}

func buildRevokeRolesButtons(product string, rolesForProduct []models.Role, account models.Account) []map[string]interface{} {
	buttons := []map[string]interface{}{}
	for _, role := range rolesForProduct {
		buttons = append(buttons, map[string]interface{}{
			"type":        "message",
			"text":        fmt.Sprintf("Revoke %v", role.Role),
			"postback_id": "send_message",
			"value":       fmt.Sprintf("%v:%v:%v:%v", RevokeRole, product, account.AccountId, role.Role),
			"user_ids":    []interface{}{},
		})
	}
	return buttons
}
