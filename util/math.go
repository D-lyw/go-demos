package util

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

func init() {
	println("test init")
}
func IsSameOfBigger(a *big.Int, b *big.Int) bool {
	if a.Cmp(b) >= 0 {
		return true
	}
	return false
}

func MulVal(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func TransferEther(num int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(num), big.NewInt(params.Ether))
}
