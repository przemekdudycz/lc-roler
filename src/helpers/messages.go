package helpers

import (
	"fmt"

	"livechat.com/lc-roler/models"
)

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
							"text":        "Agent List",
							"postback_id": "send_message",
							"value":       GetAgentsList,
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

func AgentsListMessage(chatId string, agentsList []models.Agent) models.ChatMessage {
	message := models.ChatMessage{
		ChatId: chatId,
		Event: models.ChatEvent{
			Type:       "rich_message",
			TemplateId: "cards",
			Elements:   []map[string]interface{}{},
		},
	}

	for _, agent := range agentsList {
		message.Event.Elements = append(message.Event.Elements, map[string]interface{}{
			"title":    fmt.Sprintf("Name: %v", agent.Name),
			"subtitle": fmt.Sprintf("Email: %v \n Role: %v", agent.Id, agent.Role),
			"buttons": []map[string]interface{}{
				{
					"type":        "message",
					"text":        UseViceOwnerRole,
					"postback_id": "send_message",
					"value":       fmt.Sprintf("%v:%v", agent.Id, UseViceOwnerRole),
					"user_ids":    []interface{}{},
				},
				{
					"type":        "message",
					"text":        UseAdministratorRole,
					"postback_id": "send_message",
					"value":       fmt.Sprintf("%v:%v", agent.Id, UseAdministratorRole),
					"user_ids":    []interface{}{},
				},
				{
					"type":        "message",
					"text":        UseNormalRole,
					"postback_id": "send_message",
					"value":       fmt.Sprintf("%v:%v", agent.Id, UseNormalRole),
					"user_ids":    []interface{}{},
				},
			},
		})
	}
	return message
}
