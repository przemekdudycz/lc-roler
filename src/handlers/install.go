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
)

func HandleInstall(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	code := queryParams["code"]
	accessToken := getAccessToken(strings.Join(code, ""))
	fmt.Printf("AccessToken: %v", accessToken)
	fmt.Fprintln(w, "Installation completed")
}

func getAccessToken(code string) string {
	authConfiguration := config.GetAuthConfiguration()

	reqBodyValues := map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"client_id":     authConfiguration.ClientId,
		"client_secret": authConfiguration.ClientSecret,
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

	return authResponseObject.AccessToken
}
