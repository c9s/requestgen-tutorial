package binanceapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/c9s/bbgo/pkg/testutil"
)

func getTestClientOrSkip(t *testing.T) *RestClient {
	key, secret, ok := testutil.IntegrationTestConfigured(t, "BINANCE")
	if !ok {
		t.SkipNow()
		return nil
	}

	client := NewClient(RestBaseURL)
	client.Auth(key, secret)
	return client
}

func TestClient_setTimeOffsetFromServer(t *testing.T) {
	client := NewClient(RestBaseURL)
	err := client.SetTimeOffsetFromServer(context.Background())
	assert.NoError(t, err)
}
