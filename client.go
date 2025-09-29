package landscape

import "net/http"

type LandscapeAPIClient struct {
	accountName string
	rootURL     string
	accessKey   string
	secretKey   string
	http        *http.Client
}

func NewClient(accountName, rootURL, accessKey, secretKey string) *LandscapeAPIClient {
	return &LandscapeAPIClient{
		accountName: accountName,
		rootURL:     rootURL,
		accessKey:   accessKey,
		secretKey:   secretKey,
		http:        &http.Client{},
	}
}
