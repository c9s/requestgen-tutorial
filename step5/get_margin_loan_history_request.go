package binanceapi

import (
	"encoding/json"
	"time"

	"github.com/c9s/requestgen"

	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/c9s/bbgo/pkg/types"
)

type RowsResponse struct {
	Rows  json.RawMessage `json:"rows"`
	Total int             `json:"total"`
}

// LoanStatus is one of PENDING (pending execution), CONFIRMED (successfully loaned), FAILED (execution failed, nothing happened to your account);
type LoanStatus string

const (
	LoanStatusPending   LoanStatus = "PENDING"
	LoanStatusConfirmed LoanStatus = "CONFIRMED"
	LoanStatusFailed    LoanStatus = "FAILED"
)

type MarginLoanRecord struct {
	IsolatedSymbol string                     `json:"isolatedSymbol"`
	TxId           int64                      `json:"txId"`
	Asset          string                     `json:"asset"`
	Principal      fixedpoint.Value           `json:"principal"`
	Timestamp      types.MillisecondTimestamp `json:"timestamp"`
	Status         LoanStatus                 `json:"status"`
}

// GetMarginLoanHistoryRequest
// See https://binance-docs.github.io/apidocs/spot/en/#query-loan-record-user_data
//
// The API response the following structure:
// {
//  "rows": [
//    {
//        "isolatedSymbol": "BNBUSDT", // isolated symbol, will not be returned for crossed margin
//        "txId": 12807067523,
//        "asset": "BNB",
//        "principal": "0.84624403",
//        "timestamp": 1555056425000,
//        "status": "CONFIRMED"   //one of PENDING (pending execution), CONFIRMED (successfully loaned), FAILED (execution failed, nothing happened to your account);
//    }
//  ],
//  "total": 1
// }
//
// txId or startTime must be sent. txId takes precedence.
// Response in descending order
// If isolatedSymbol is not sent, crossed margin data will be returned
// The max interval between startTime and endTime is 30 days.
// If startTime and endTime not sent, return records of the last 7 days by default
// Set archived to true to query data from 6 months ago
//
//go:generate requestgen -method GET -url "/sapi/v1/margin/loan" -type GetMarginLoanHistoryRequest -responseType .RowsResponse -responseDataField Rows -responseDataType []MarginLoanRecord
type GetMarginLoanHistoryRequest struct {
	client requestgen.AuthenticatedAPIClient

	asset          string     `param:"asset"`
	startTime      *time.Time `param:"startTime,milliseconds"`
	endTime        *time.Time `param:"endTime,milliseconds"`
	isolatedSymbol *string    `param:"isolatedSymbol"`
	archived       *bool      `param:"archived"`
	size           *int       `param:"size"`
	current        *int       `param:"current"`
}

func (c *RestClient) NewGetMarginLoanHistoryRequest() *GetMarginLoanHistoryRequest {
	return &GetMarginLoanHistoryRequest{client: c}
}
