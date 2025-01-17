package helpers

import "os"

type EnvVar string

var EnvVars = struct {
	AccountId      EnvVar
	Endpoint       EnvVar
	PlatformApiKey EnvVar
	TfLog          EnvVar
}{
	AccountId:      "HARNESS_ACCOUNT_ID",
	Endpoint:       "HARNESS_ENDPOINT",
	PlatformApiKey: "HARNESS_PLATFORM_API_KEY",
	TfLog:          "TF_LOG",
}

func (e EnvVar) String() string {
	return string(e)
}

func (e EnvVar) Get() string {
	return os.Getenv(string(e))
}

func (e EnvVar) GetWithDefault(fallback string) string {
	if value, ok := os.LookupEnv(string(e)); ok {
		return value
	}
	return fallback
}
