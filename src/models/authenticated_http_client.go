package models

import (
	"fmt"
	"net/http"
	"time"
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

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &AuthenticatedHttpClient{
		httpClient: httpClient,
	}
}

func (ac *AuthenticatedHttpClient) Init(accessToken *string, refreshToken *string) {
	authenticatedClient.accessToken = accessToken
	authenticatedClient.refreshToken = refreshToken
}

func (ac *AuthenticatedHttpClient) PerformRequest(r *http.Request) (*http.Response, error) {
	if authenticatedClient.accessToken == nil {
		return nil, fmt.Errorf("you need to authenticate the application")
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", authenticatedClient.accessToken))
	return authenticatedClient.httpClient.Do(r)
}
