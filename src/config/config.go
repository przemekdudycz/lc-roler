package config

// Auth
type AuthConfiguration struct {
	ClientId     string
	ClientSecret string
}

var authConfiguration = AuthConfiguration{
	ClientId:     "*",
	ClientSecret: "*",
}

func GetAuthConfiguration() (a AuthConfiguration) {
	return authConfiguration
}
