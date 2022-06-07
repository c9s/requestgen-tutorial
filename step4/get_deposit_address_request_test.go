package binanceapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestClient_GetDepositAddressRequest(t *testing.T) {
	client := getTestClientOrSkip(t)
	ctx := context.Background()

	err := client.SetTimeOffsetFromServer(ctx)
	assert.NoError(t, err)

	req := client.NewGetDepositAddressRequest()
	req.Coin("BTC")
	address, err := req.Do(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, address)
	assert.NotEmpty(t, address.Url)
	assert.NotEmpty(t, address.Address)
	t.Logf("deposit address: %+v", address)
}


