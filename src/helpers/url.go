package helpers

import "strings"

var (
	integrationUrl = "http://edca88d66850.ngrok.io/install"
	// External URLs
	authBaseUrl      = "https://accounts.labs.livechat.com"
	authApiVersion   = "v2"
	configBaseUrl    = "https://api.labs.livechatinc.com"
	configApiVersion = "v3.3"
)

func IntegrationUrl() string {
	return integrationUrl
}

func AccessTokenUrl() string {
	return strings.Join([]string{authBaseUrl, authApiVersion, "token"}, "/")
}

func WebhooksListUrl() string {
	return strings.Join([]string{configBaseUrl, configApiVersion, "configuration", "action", "list_webhooks"}, "/")
}

func RegisterWebhookUrl() string {
	return strings.Join([]string{configBaseUrl, configApiVersion, "configuration", "action", "register_webhook"}, "/")
}

func EnableWebhookUrl() string {
	return strings.Join([]string{configBaseUrl, configApiVersion, "configuration", "action", "enable_license_webhooks"}, "/")
}
