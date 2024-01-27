package listen

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func GetERC721TransferEventQuery() *ethereum.FilterQuery {
	return &ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), // USDT in Ethereum
		},
		Topics: [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}
}

func GetERC721ApprovalEventQuery() {

}
