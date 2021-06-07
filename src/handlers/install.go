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
	authenticatedHttpClient.Init(strings.Join(code, ""), authConfiguration.ClientId, authConfiguration.ClientSecret)

	existingWebhooks := models.GetWebhooksList(authConfiguration.ClientId)
	fmt.Printf("Existing webhooks: %v", existingWebhooks)

	// isNewChatWebhookExists := models.IsWebhookWithActionInSlice(existingWebhooks, "incoming_chat")
	// if !isNewChatWebhookExists {
	// 	// registerWebhook()
	// 	// enableWebhook()
	// }

	fmt.Fprintln(w, "Installation completed")
}
