package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func (ac *AuthenticatedHttpClient) SendAuthenticatedRequest(r *http.Request, accessSource string) (*http.Response, error) {
	if ac.authData[accessSource].AccessToken == nil {
		return nil, fmt.Errorf("you need to authenticate the application")
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", *ac.authData[accessSource].AccessToken))
	return ac.httpClient.Do(r)
}

func (ac *AuthenticatedHttpClient) SendRequest(req *http.Request, accessSource string, returnObject interface{}) error {
	res, err := ac.SendAuthenticatedRequest(req, accessSource)

	fmt.Printf("HTTP response status code: %v \n", res.StatusCode)

	if err != nil {
		fmt.Printf("there was HTTP call error: %v", err.Error())
		return err
	}
	defer res.Body.Close()

	if returnObject != nil {
		bodyBytes, _ := io.ReadAll(res.Body)
		jsonError := json.Unmarshal(bodyBytes, returnObject)

		if jsonError != nil {
			fmt.Printf("there was Unmarshal error: %v", jsonError.Error())
			return jsonError
		}
	}

	return nil
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

	authResponseObject := struct {
		AccessToken string `json:"access_token"`
	}{}
	err := httpClient.SendRequest(req, "agentAuth", &authResponseObject)

	if err != nil {
		fmt.Printf("there was an error when GetCustomerAccessTokens: %v", err.Error())
	}

	return authResponseObject.AccessToken
}
