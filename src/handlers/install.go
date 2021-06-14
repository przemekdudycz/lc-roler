package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"livechat.com/lc-roler/config"
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
	rmPostbackWebhookAction := "incoming_rich_message_postback"
	eventWebhookAction := "incoming_event"

	destinationNewChatWebhookUrl := config.DestinationNewChatWebhookUrl()
	destinationRMPostbackWebhookUrl := config.DestinationRMPostbackWebhookUrl()
	destintionEventWebhookUrl := config.DestinationEventWebhookUrl()

	isNewChatWebhookExists := models.IsWebhookWithActionInSlice(existingWebhooks, newChatWebhookAction)
	isRMPostbackWebhookExists := models.IsWebhookWithActionInSlice(existingWebhooks, rmPostbackWebhookAction)
	isEventWebhookExists := models.IsWebhookWithActionInSlice(existingWebhooks, eventWebhookAction)

	if !isNewChatWebhookExists {
		webhookData := models.WebhookRequestBody{
			Url:           destinationNewChatWebhookUrl,
			Description:   "New chat lc-roler webhook",
			Action:        newChatWebhookAction,
			SecretKey:     "verysecretkey",
			OwnerClientId: authConfiguration.ClientId,
			Type:          "license",
		}
		webhookId := models.RegisterWebhook(webhookData)
		fmt.Printf("NewChatWebhookId: %v \n", webhookId)
	}

	if !isRMPostbackWebhookExists {
		webhookData := models.WebhookRequestBody{
			Url:           destinationRMPostbackWebhookUrl,
			Description:   "New chat lc-roler webhook",
			Action:        rmPostbackWebhookAction,
			SecretKey:     "verysecretkey",
			OwnerClientId: authConfiguration.ClientId,
			Type:          "license",
		}
		webhookId := models.RegisterWebhook(webhookData)
		fmt.Printf("NewEventWebhookId: %v \n", webhookId)
	}

	if !isEventWebhookExists {
		webhookData := models.WebhookRequestBody{
			Url:           destintionEventWebhookUrl,
			Description:   "New chat lc-roler webhook",
			Action:        eventWebhookAction,
			SecretKey:     "verysecretkey",
			OwnerClientId: authConfiguration.ClientId,
			Type:          "license",
			Filters:       models.WebhookFilter{AuthorType: "customer"},
		}
		webhookId := models.RegisterWebhook(webhookData)
		fmt.Printf("NewEventWebhookId: %v \n", webhookId)
	}

	if !isNewChatWebhookExists || !isRMPostbackWebhookExists || isEventWebhookExists {
		models.EnableWebhooks()
	}

	fmt.Fprintln(w, "Installation completed")
}
