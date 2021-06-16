package models

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	accountsListResponse := []Account{}
	err := httpClient.SendRequest(req, "agentAuth", &accountsListResponse)

	if err != nil {
		fmt.Printf("there was an error when GetAccountsList: %v", err)
	}

	return accountsListResponse
}

func GetAccount(accountId string) Account {
	httpClient := GetAuthenticatedHttpClient()
	getAccountUrl := config.GetAccountUrl(accountId)

	req, _ := http.NewRequest("GET", getAccountUrl, nil)
	accountResponse := Account{}
	err := httpClient.SendRequest(req, "agentAuth", &accountResponse)

	if err != nil {
		fmt.Printf("there was an error when GetAccount: %v", err)
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

	accountResponse := Account{}
	err := httpClient.SendRequest(req, "agentAuth", &accountResponse)

	if err != nil {
		fmt.Printf("there was an error when UpdateAccountRoles: %v", err)
	}

	return accountResponse
}
