package helpers

func WelcomeMessage(chatId string) map[string]interface{} {
	return map[string]interface{}{
		"chat_id": chatId,
		"event": map[string]interface{}{
			"type":        "rich_message",
			"template_id": "quick_replies",
			"elements": []map[string]interface{}{
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
