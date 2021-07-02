package main

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func blockTimeStamp(hash common.Hash) (stamp uint64, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	head, err := client.HeaderByHash(ctx, hash)
	if err == nil {
		stamp = head.Time
	}
	return
}
