package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"livechat.com/lc-roler/config"
)

type Role struct {
	RoleId     string `json:"role_id"`
	Product    string `json:"product"`
	Role       string `json:"role"`
	Type       string `json:"type"`
	Predefined bool   `json:"predefined"`
}

type Account struct {
	AccountId string `json:"account_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Roles     []Role `json:"roles"`
}

func GetAccountsList() []Account {
	httpClient := GetAuthenticatedHttpClient()
	getAccountList := config.GetAccountsListUrl()

	req, _ := http.NewRequest("GET", getAccountList, nil)

	res, err := httpClient.PerformRequest(req, "agentAuth")

	fmt.Printf("GetAccountsListStatusCode: %v", res.StatusCode)

	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("Invalid GetAccountsList response: %v", res.StatusCode)
	}

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	accountsListResponse := []Account{}
	jsonError := json.Unmarshal(bodyBytes, &accountsListResponse)

	if jsonError != nil {
		panic(err)
	}

	return accountsListResponse
}

func GetAccount(accountId string) Account {
	httpClient := GetAuthenticatedHttpClient()
	getAccountUrl := config.GetAccountUrl(accountId)

	req, _ := http.NewRequest("GET", getAccountUrl, nil)

	res, err := httpClient.PerformRequest(req, "agentAuth")

	fmt.Printf("GetAccountStatusCode: %v", res.StatusCode)

	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("Invalid GetAccount response: %v", res.StatusCode)
	}

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	accountResponse := Account{}
	jsonError := json.Unmarshal(bodyBytes, &accountResponse)

	if jsonError != nil {
		panic(err)
	}

	return accountResponse
}

func UpdateAccountRoles(accountId string, setRoles []Role, deleteRoles []Role) Account {
	httpClient := GetAuthenticatedHttpClient()
	updateAccountRolesUrl := config.UpdateAccountRolesUrl(accountId)

	reqBody := map[string][]Role{
		"set_roles":    setRoles,
		"delete_roles": deleteRoles,
	}
	reqJsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("PUT", updateAccountRolesUrl, bytes.NewReader(reqJsonBody))

	res, err := httpClient.PerformRequest(req, "agentAuth")

	fmt.Printf("UpdateAccountsRole status code: %v", res.StatusCode)

	if err != nil {
		fmt.Printf("There was an error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		fmt.Printf("Invalid UpdateAccountsRole response: %v", res.StatusCode)
	}

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	accountResponse := Account{}
	jsonError := json.Unmarshal(bodyBytes, &accountResponse)

	if jsonError != nil {
		panic(err)
	}

	return accountResponse
}
