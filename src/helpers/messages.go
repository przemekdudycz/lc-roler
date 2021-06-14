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
							"text":        GetAgentsList,
							"postback_id": "send_message",
							"value":       GetAgentsList,
							"user_ids":    []interface{}{},
						},
						{
							"type":        "message",
							"text":        GetRolesList,
							"postback_id": "send_message",
							"value":       GetRolesList,
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
		})
	}
	return message
}
