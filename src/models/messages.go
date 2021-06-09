package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"livechat.com/lc-roler/config"
	"livechat.com/lc-roler/helpers"
)

func SendMessageToCustomer(chatId string) string {
	httpClient := GetAuthenticatedHttpClient()
	config := config.GetAuthConfiguration()
	sendEventUrl := helpers.GetSendEventUrl(config.LicenseId)

	reqBodyValues := map[string]interface{}{
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
							"text":        "GetAgentsList",
							"postback_id": "send_message",
							"value":       "agentsList",
							"user_ids":    []interface{}{},
						},
					},
				},
			},
		},
	}
	reqBody, _ := json.Marshal(reqBodyValues)
	req, _ := http.NewRequest("POST", sendEventUrl, bytes.NewReader(reqBody))

	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.PerformRequest(req, "agentAuth")
	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("Send message error: %v \n", res.StatusCode)
	}

	defer res.Body.Close()
	raw, _ := ioutil.ReadAll(res.Body)

	return string(raw)
}
