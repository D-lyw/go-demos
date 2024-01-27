package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"log"
	"math/big"
)

func QueryTxBlock() {
	// 查询区块头信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number, header.Hash())

	//	查询区块信息
	block, err := client.BlockByNumber(context.Background(), big.NewInt(17735180))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Hash(), block.TxHash(), block.Size(), block.GasUsed(), block.Difficulty(), len(block.Transactions()))

	// 遍历区块里面的每一个 tx
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash(), tx.To(), tx.Nonce(), tx.Gas())
	}
}

// 转账交易
func TxTransfer() {

	//privateKey, err := crypto.HexToECDSA("")

	memonic := "wall load element exit fetch mistake elephant absent sail motor excite drum"

	seed := bip39.NewSeed(memonic, "")
	masterKey, _ := bip32.NewMasterKey(seed)
	fmt.Println(masterKey.PublicKey().String())

	fmt.Println(seed, hexutil.Encode(seed))

	//walllet, err := hdwallet.NewFromMnemonic(memonic)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//path := hdwallet.MustParseDerivationPath("m/44'60'/0'/0/0")
	//account, err := walllet.Derive(path, true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(account.Address.Hex())
}
