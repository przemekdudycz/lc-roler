package config

// Auth
type AuthConfiguration struct {
	ClientId     string
	ClientSecret string
	LicenseId    string
}

var authConfiguration = AuthConfiguration{
	ClientId:     "*",
	ClientSecret: "*",
	LicenseId:    "100101160",
}

func GetAuthConfiguration() (a AuthConfiguration) {
	return authConfiguration
}
