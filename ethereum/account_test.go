package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func TestWalletGenerate(t *testing.T) {

	WalletGenerate()

}

func TestAccountBalance(t *testing.T) {

	AccountBalance()

}

func TestNone(t *testing.T) {
	value := big.NewInt(1).Mul(big.NewInt(5000000), big.NewInt(params.GWei))
	fmt.Println(value)
}

func TestTransfer(t *testing.T) {
	Transfer()
}
