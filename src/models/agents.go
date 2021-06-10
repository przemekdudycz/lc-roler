package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"livechat.com/lc-roler/helpers"
)

type Agent struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func GetAgentsList() []Agent {
	httpClient := GetAuthenticatedHttpClient()
	getAgentsListUrl := helpers.GetAgentsListUrl()

	reqBody, _ := json.Marshal(map[string]interface{}{})

	req, _ := http.NewRequest("POST", getAgentsListUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.PerformRequest(req, "agentAuth")
	fmt.Printf("GetAgentsListStatusCode: %v", res.StatusCode)

	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("Invalid GetAgentList response: %v", res.StatusCode)
	}

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	agentResponse := []Agent{}
	jsonError := json.Unmarshal(bodyBytes, &agentResponse)

	if jsonError != nil {
		panic(err)
	}

	return agentResponse
}
