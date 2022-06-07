package binanceapi

import (
	"time"

	"github.com/c9s/requestgen"

	"github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/c9s/bbgo/pkg/types"
)

type DepositHistory struct {
	Amount        fixedpoint.Value           `json:"amount"`
	Coin          string                     `json:"coin"`
	Network       string                     `json:"network"`
	Status        int                        `json:"status"`
	Address       string                     `json:"address"`
	AddressTag    string                     `json:"addressTag"`
	TxId          string                     `json:"txId"`
	InsertTime    types.MillisecondTimestamp `json:"insertTime"`
	TransferType  int                        `json:"transferType"`
	UnlockConfirm int                        `json:"unlockConfirm"`
	ConfirmTimes  string                     `json:"confirmTimes"`
}

// GetDepositHistoryRequest returns the following structure:
// [
//    {
//        "amount":"0.00999800",
//        "coin":"PAXG",
//        "network":"ETH",
//        "status":1,
//        "address":"0x788cabe9236ce061e5a892e1a59395a81fc8d62c",
//        "addressTag":"",
//        "txId":"0xaad4654a3234aa6118af9b4b335f5ae81c360b2394721c019b5d1e75328b09f3",
//        "insertTime":1599621997000,
//        "transferType":0,
//        "unlockConfirm":"12/12",  // confirm times for unlocking
//        "confirmTimes":"12/12"
//    },
//    {
//        "amount":"0.50000000",
//        "coin":"IOTA",
//        "network":"IOTA",
//        "status":1,
//        "address":"SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW",
//        "addressTag":"",
//        "txId":"ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999",
//        "insertTime":1599620082000,
//        "transferType":0,
//        "unlockConfirm":"1/12",
//        "confirmTimes":"1/1"
//    }
// ]
//go:generate requestgen -method GET -url "/sapi/v1/capital/deposit/hisrec" -type GetDepositHistoryRequest -responseType []DepositHistory
type GetDepositHistoryRequest struct {
	client requestgen.AuthenticatedAPIClient

	coin *string `param:"coin"`

	startTime *time.Time `param:"startTime,milliseconds"`
	endTime   *time.Time `param:"endTime,milliseconds"`
}

func (c *RestClient) NewGetDepositHistoryRequest() *GetDepositHistoryRequest {
	return &GetDepositHistoryRequest{client: c}
}
