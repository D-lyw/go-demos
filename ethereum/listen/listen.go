package listen

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"os"
)

/// 持续监听链上指定合约事件
//1. 连接区块链节点
//2. 订阅对应合约事件
//3. 处理对应事件逻辑

func NewEthClient() *ethclient.Client {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	rpcKey := os.Getenv("RPC_KEY")

	client, err := ethclient.Dial(fmt.Sprintf("wss://eth-mainnet.g.alchemy.com/v2/%s", rpcKey))
	if err != nil {
		panic(err)
	}

	return client
}

func HandleListenFilter(client *ethclient.Client, query *ethereum.FilterQuery) {
	log := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), *query, log)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			fmt.Print(err)
		case vlog := <-log:
			fmt.Println("From: ", common.HexToAddress(vlog.Topics[1].String()))
			fmt.Println("To: ", common.HexToAddress(vlog.Topics[2].Hex()))
			//fmt.Println("TokenId: ", vlog.Topics[3].Big().String())
			fmt.Println(vlog)
		}
	}
}
