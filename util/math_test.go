package util

import (
	"math/big"
	"testing"
)

func TestIsSameOfBigger(t *testing.T) {
	a := big.NewInt(9424289238)
	b, _ := new(big.Int).SetString("9424289238", 10)
	if !IsSameOfBigger(a, b) {
		t.Error("Compare Error")
	}
}

func TestMulVal(t *testing.T) {
	a := big.NewInt(10)
	b := big.NewInt(20)
	result := big.NewInt(200)
	if MulVal(a, b).Cmp(result) != 0 {
		t.Error("Mul value logic error")
	}
}

func TestTransferEther(t *testing.T) {
	var a int64 = 1
	println(TransferEther(a))
}
