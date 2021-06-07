package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"livechat.com/lc-roler/config"
	"livechat.com/lc-roler/helpers"
	"livechat.com/lc-roler/models"
)

func HandleInstall(w http.ResponseWriter, r *http.Request) {
	authConfiguration := config.GetAuthConfiguration()
	queryParams := r.URL.Query()
	code := queryParams["code"]
	accessToken, refreshToken := getAccessToken(strings.Join(code, ""), authConfiguration.ClientId, authConfiguration.ClientSecret)
	fmt.Printf("AccessToken: %v", accessToken)

	authenticatedHttpClient := models.GetAuthenticatedHttpClient()
	authenticatedHttpClient.Init(&accessToken, &refreshToken)

	getWebhooksList(authConfiguration.ClientId)

	fmt.Fprintln(w, "Installation completed")
}

func getAccessToken(code string, clientId string, clientSecret string) (string, string) {
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

func getWebhooksList(clientId string) {
	getWebhooksListUrl := helpers.WebhooksListUrl()
	httpClient := models.GetAuthenticatedHttpClient()

	reqBodyValues := map[string]interface{}{
		"owner_client_id": clientId,
	}
	reqBody, _ := json.Marshal(reqBodyValues)
	req, _ := http.NewRequest("POST", getWebhooksListUrl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	// resPayload := make(map[string]map[string]interface{})
	res, _ := httpClient.PerformRequest(req)

	defer res.Body.Close()

	raw, _ := ioutil.ReadAll(res.Body)

	// err := json.Unmarshal(raw, &resPayload)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("WebhooksList response: %v", string(raw))
}
