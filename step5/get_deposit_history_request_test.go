package binanceapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetDepositHistoryRequest(t *testing.T) {
	client := getTestClientOrSkip(t)
	ctx := context.Background()

	err := client.SetTimeOffsetFromServer(ctx)
	assert.NoError(t, err)

	req := client.NewGetDepositHistoryRequest()
	history, err := req.Do(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, history)
	assert.NotEmpty(t, history)
	t.Logf("deposit history: %+v", history)
}
