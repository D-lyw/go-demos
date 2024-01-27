package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://mainnet.infura.io/v3/2e1e3366832e41368179dc9e08156d85")
	if err != nil {
		log.Fatal("Init client error: ", err)
	}
}

func AccountBalance() {

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress("0x431B4CA18E269Fc7e1F5AF49B9F4E2AF683f6207"), nil)
	if err != nil {
		log.Fatal("Query balance error: ", err)
	}
	fmt.Println(balance, balance.String())
}

// https://goethereumbook.org/zh/wallet-generate/
func WalletGenerate() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKey)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("PrivateKey: ", hexutil.Encode(privateKeyBytes)[2:])

	// 生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("PublicKey: ", hexutil.Encode(publicKeyBytes))

	// 生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address: ", address)

	// 地址生成逻辑：将公钥进行 Keccak-256 哈希计算，取最后 20 个字节
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("Hashed Address: ", hexutil.Encode(hash.Sum(nil)[12:]))
}

// Transfer ETH 转账交易 https://goethereumbook.org/zh/transfer-eth/
func Transfer() {

	goerliClient, err := ethclient.Dial("https://goerli.infura.io/v3/2e1e3366832e41368179dc9e08156d85")
	if err != nil {
		log.Fatal(err)
	}

	// 开发测试账号
	privateKey, err := crypto.HexToECDSA("b68886b709b09cec1ae2144a9985f112f74269bb4cb800ab668ae55eed2b563f")

	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 构造交易参数
	nonce, err := goerliClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(21000)
	value := new(big.Int).Mul(big.NewInt(5000000), big.NewInt(params.GWei))
	gasprice, err := goerliClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x66567071D55A9FBE6B3944172592961c1C414075")

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		Gas:      gasLimit,
		GasPrice: gasprice,
		To:       &toAddress,
		Value:    value,
		Data:     nil,
	})

	chainId, err := goerliClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = goerliClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tx send, hash is: ", signedTx.Hash())
}
