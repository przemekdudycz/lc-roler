package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"livechat.com/lc-roler/helpers"
)

type AuthenticatedHttpClient struct {
	httpClient   *http.Client
	accessToken  *string
	refreshToken *string
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
	}
	return authenticatedClient
}

func (ac *AuthenticatedHttpClient) Init(code string, clientId string, clientSecret string) {
	accessToken, refreshToken := getAuthTokens(code, clientId, clientSecret)
	fmt.Printf("AccessToken: %v \n", accessToken)
	fmt.Printf("RefreshToken: %v \n", refreshToken)
	ac.accessToken = &accessToken
	ac.refreshToken = &refreshToken
}

func (ac *AuthenticatedHttpClient) PerformRequest(r *http.Request) (*http.Response, error) {
	if ac.accessToken == nil {
		return nil, fmt.Errorf("you need to authenticate the application")
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", *ac.accessToken))
	return ac.httpClient.Do(r)
}

func getAuthTokens(code string, clientId string, clientSecret string) (string, string) {
	reqBodyValues := map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"client_id":     clientId,
		"client_secret": clientSecret,
		"redirect_uri":  helpers.IntegrationUrl(),
	}
	jsonReqBody, _ := json.Marshal(reqBodyValues)

	authResponse, _ := http.Post(helpers.AccessTokenUrl(), "application/json", bytes.NewBuffer(jsonReqBody))
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
