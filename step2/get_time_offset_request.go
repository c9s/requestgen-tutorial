package binanceapi

import (
	"github.com/c9s/bbgo/pkg/types"
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/api/v3/time" -type GetTimeOffsetRequest -responseType .TimeOffsetResponse
type GetTimeOffsetRequest struct {
	client requestgen.APIClient
}

type TimeOffsetResponse struct {
	ServerTime types.MillisecondTimestamp `json:"serverTime"`
}

