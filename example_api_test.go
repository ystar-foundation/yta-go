package yta_test

import "os"

func getAPIURL() string {
	apiURL := os.Getenv("EOS_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://mainnet.ystar-foundation.com"
}
