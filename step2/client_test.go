package binanceapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_setTimeOffsetFromServer(t *testing.T) {
	client := NewClient("")
	err := client.SetTimeOffsetFromServer(context.Background())
	assert.NoError(t, err)
}
