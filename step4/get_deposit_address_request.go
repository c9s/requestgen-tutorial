package binanceapi

import (
	"github.com/c9s/requestgen"
)

type DepositAddress struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

// GetDepositAddressRequest returns the address of deposit:
// If network is not send, return with default network of the coin.
// You can get network and isDefault in networkList in the response of Get /sapi/v1/capital/config/getall (HMAC SHA256).
//
// {
//    "address": "1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv",
//    "coin": "BTC",
//    "tag": "",
//    "url": "https://btc.com/1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv"
// }
//
//go:generate requestgen -method GET -url "/sapi/v1/capital/deposit/address" -type GetDepositAddressRequest -responseType .DepositAddress
type GetDepositAddressRequest struct {
	client requestgen.AuthenticatedAPIClient

	coin string `param:"coin"`

	network *string `param:"network"`
}

func (c *RestClient) NewGetDepositAddressRequest() *GetDepositAddressRequest {
	return &GetDepositAddressRequest{client: c}
}
