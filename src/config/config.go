package config

// Auth
type AuthConfiguration struct {
	ClientId     string
	ClientSecret string
	LicenseId    string
}

var authConfiguration = AuthConfiguration{
	ClientId:     "edfb1c72783622d68fa8e93ab9ac362a",
	ClientSecret: "952a1b9be91169348681796fd97180f3550fd962",
	LicenseId:    "100101160",
}

func GetAuthConfiguration() (a AuthConfiguration) {
	return authConfiguration
}
