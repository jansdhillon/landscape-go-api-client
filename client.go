package landscape

import "net/http"

type LandscapeAPIClient struct {
	AccountName string
	RootURL     string
	AccessKey   string
	SecretKey   string
	HTTP        *http.Client
}

func NewLandscapeAPIClient(accountName, rootURL, accessKey, secretKey string) *LandscapeAPIClient {
	return &LandscapeAPIClient{
		AccountName: accountName,
		RootURL:     rootURL,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		HTTP:        &http.Client{},
	}
}
