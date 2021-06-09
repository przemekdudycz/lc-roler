package handlers

import (
	"fmt"
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

	authenticatedHttpClient := models.GetAuthenticatedHttpClient()
	accessToken, refreshToken := models.GetAuthTokens(strings.Join(code, ""), authConfiguration.ClientId, authConfiguration.ClientSecret)
	fmt.Printf("AccessToken: %v \n", accessToken)
	fmt.Printf("RefreshToken: %v \n", refreshToken)
	authenticatedHttpClient.AddAuthData(&models.AuthData{AccessToken: &accessToken, RefreshToken: &refreshToken}, "agentAuth")

	customerAccessToken := models.GetCustomerAccessTokens(authConfiguration.ClientId)
	fmt.Printf("Customer access token: %v \n", customerAccessToken)
	authenticatedHttpClient.AddAuthData(&models.AuthData{AccessToken: &customerAccessToken}, "customerAuth")

	existingWebhooks := models.GetWebhooksList(authConfiguration.ClientId)
	fmt.Printf("Existing webhooks: %v", existingWebhooks)

	newChatWebhookAction := "incoming_chat"
	destinationNewChatWebhookUrl := helpers.DestinationNewChatWebhookUrl()
	isNewChatWebhookExists := models.IsWebhookWithActionInSlice(existingWebhooks, newChatWebhookAction)
	if !isNewChatWebhookExists {
		webhookId := models.RegisterWebhook(authConfiguration.ClientId, newChatWebhookAction, destinationNewChatWebhookUrl)
		fmt.Printf("WebhookId: %v \n", webhookId)
		models.EnableWebhooks()
	}

	fmt.Fprintln(w, "Installation completed")
}
