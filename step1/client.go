package binanceapi

import (
	"net/http"
	"net/url"
	"time"

	"github.com/c9s/requestgen"
)

const defaultHTTPTimeout = time.Second * 15
const RestBaseURL = "https://api.binance.com"

var DefaultHttpClient = &http.Client{
	Timeout: defaultHTTPTimeout,
}

type RestClient struct {
	requestgen.BaseAPIClient

	Key, Secret string

	recvWindow int
	timeOffset int64
}

func NewClient(baseURL string) *RestClient {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	client := &RestClient{
		BaseAPIClient: requestgen.BaseAPIClient{
			BaseURL:    u,
			HttpClient: DefaultHttpClient,
		},
	}

	return client
}

func (c *RestClient) Auth(key, secret string) {
	c.Key = key
	c.Secret = secret
}
