package binanceapi

import (
	"context"
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

func (c *RestClient) SetTimeOffsetFromServer(ctx context.Context) error {
	req := &GetTimeOffsetRequest{client: c}
	resp, err := req.Do(ctx)
	if err != nil {
		return err
	}

	c.timeOffset = currentTimestamp() - resp.ServerTime.Time().UnixMilli()
	return nil
}

func currentTimestamp() int64 {
	return time.Now().UnixMilli()
}
