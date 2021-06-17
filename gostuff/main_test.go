package main

import (
	"fmt"
	"testing"

	"github.com/DaveAppleton/ether_go/ethKeys"
)

func TestAll(t *testing.T) {
	mainnet := ethKeys.NewKey("../.mainnet")
	mainnet.RestoreOrCreate()
	fmt.Println(mainnet.PublicKeyAsHexString())
	t.Fail()
}
