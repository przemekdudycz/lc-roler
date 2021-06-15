package config

import (
	"fmt"
	"strings"
)

var (
	baseAppUrl = "http://2e49257e647b.ngrok.io"
	// External URLs
	accountsBaseUrl    = "https://accounts.labs.livechat.com"
	accountsApiVersion = "v2"
	configBaseUrl      = "https://api.labs.livechatinc.com"
	configApiVersion   = "v3.3"
	customerApiBaseUrl = "https://api.labs.livechatinc.com"
	customerApiVersion = "v3.3"
)

func IntegrationUrl() string {
	return strings.Join([]string{baseAppUrl, "install"}, "/")
}

func AccessTokenUrl() string {
	return strings.Join([]string{accountsBaseUrl, accountsApiVersion, "token"}, "/")
}

func CustomerAccessTokenUrl() string {
	return strings.Join([]string{accountsBaseUrl, "customer"}, "/")
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

func DestinationNewChatWebhookUrl() string {
	return strings.Join([]string{baseAppUrl, "newchat"}, "/")
}

func DestinationRMPostbackWebhookUrl() string {
	return strings.Join([]string{baseAppUrl, "rmpostback"}, "/")
}

func DestinationEventWebhookUrl() string {
	return strings.Join([]string{baseAppUrl, "newevent"}, "/")
}

func GetSendEventUrl(licenseId string) string {
	url := strings.Join([]string{customerApiBaseUrl, customerApiVersion, "agent", "action", "send_event"}, "/")
	return fmt.Sprintf("%v?license_id=%v", url, licenseId)
}

func GetAgentsListUrl() string {
	return strings.Join([]string{configBaseUrl, configApiVersion, "configuration", "action", "list_agents"}, "/")
}

func UpdateAgentUrl() string {
	return strings.Join([]string{configBaseUrl, configApiVersion, "configuration", "action", "update_agent"}, "/")
}

func GetAccountsListUrl() string {
	return strings.Join([]string{accountsBaseUrl, accountsApiVersion, "accounts"}, "/")
}

func UpdateAccountRolesUrl(accountId string) string {
	return strings.Join([]string{accountsBaseUrl, accountsApiVersion, "accounts", accountId, "roles"}, "/")
}
