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

func SendMessageToCustomer(reqBodyValues map[string]interface{}) string {
	httpClient := GetAuthenticatedHttpClient()
	config := config.GetAuthConfiguration()
	sendEventUrl := helpers.GetSendEventUrl(config.LicenseId)

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
