package listen

import "github.com/ethereum/go-ethereum/ethclient"

/// 持续监听链上指定合约事件
//1. 连接区块链节点
//2. 订阅对应合约事件
//3. 处理对应事件逻辑

func NewEthClient() *ethclient.Client {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2e1e3366832e41368179dc9e08156d85")
	if err != nil {
		panic(err)
	}
	return client
}
