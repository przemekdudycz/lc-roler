package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"livechat.com/lc-roler/config"
)

type AuthData struct {
	AccessToken  *string
	RefreshToken *string
}

type AuthenticatedHttpClient struct {
	httpClient *http.Client
	authData   map[string]AuthData
}

var (
	authenticatedClient *AuthenticatedHttpClient
)

func GetAuthenticatedHttpClient() *AuthenticatedHttpClient {
	if authenticatedClient != nil {
		return authenticatedClient
	}

	fmt.Printf("Initialization of auth client \n")
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	authenticatedClient = &AuthenticatedHttpClient{
		httpClient: httpClient,
		authData:   make(map[string]AuthData),
	}
	return authenticatedClient
}

func (ac *AuthenticatedHttpClient) AddAuthData(authData *AuthData, authSource string) {
	ac.authData[authSource] = *authData
}

func (ac *AuthenticatedHttpClient) PerformRequest(r *http.Request, accessSource string) (*http.Response, error) {
	if ac.authData[accessSource].AccessToken == nil {
		return nil, fmt.Errorf("you need to authenticate the application")
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", *ac.authData[accessSource].AccessToken))
	return ac.httpClient.Do(r)
}

func GetAuthTokens(code string, clientId string, clientSecret string) (string, string) {
	reqBodyValues := map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"client_id":     clientId,
		"client_secret": clientSecret,
		"redirect_uri":  config.IntegrationUrl(),
	}
	jsonReqBody, _ := json.Marshal(reqBodyValues)

	authResponse, _ := http.Post(config.AccessTokenUrl(), "application/json", bytes.NewBuffer(jsonReqBody))
	defer authResponse.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(authResponse.Body)

	authResponseObject := struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{}
	err := json.Unmarshal(bodyBytes, &authResponseObject)

	if err != nil {
		panic(err)
	}

	return authResponseObject.AccessToken, authResponseObject.RefreshToken
}

func GetCustomerAccessTokens(clientId string) string {
	httpClient := GetAuthenticatedHttpClient()
	customerAccessTokenUrl := config.CustomerAccessTokenUrl()
	integrationUrl := config.IntegrationUrl()

	reqBodyValues := map[string]string{
		"client_id":     clientId,
		"response_type": "token",
		"redirect_uri":  integrationUrl,
	}
	reqBody, _ := json.Marshal(reqBodyValues)

	req, _ := http.NewRequest("POST", customerAccessTokenUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	res, _ := httpClient.PerformRequest(req, "agentAuth")

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("Response body: %v", string(bodyBytes))
	authResponseObject := struct {
		AccessToken string `json:"access_token"`
	}{}

	err := json.Unmarshal(bodyBytes, &authResponseObject)
	if err != nil {
		panic(err)
	}
	return authResponseObject.AccessToken
}
