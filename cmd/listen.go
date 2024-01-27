package main

import "go-demos/ethereum/listen"

func main() {
	client := listen.NewEthClient()

	filterQuery := listen.GetERC721TransferEventQuery()

	listen.HandleListenFilter(client, filterQuery)
}
