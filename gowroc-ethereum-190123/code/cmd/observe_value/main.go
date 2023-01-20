package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seblw/eth-gowroc/pkg/storage"
)

func main() {
	var (
		nodeAddr     = "ws://127.0.0.1:8545/"
		contractAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	)

	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, nodeAddr)
	if err != nil {
		log.Fatal("DialContext: ", err)
	}

	addr := common.HexToAddress(contractAddr)
	instance, err := storage.NewStorage(addr, cli)
	if err != nil {
		log.Fatal("NewStorage: ", err)
	}
	logs := make(chan *storage.StorageValueChanged)
	sub, err := instance.WatchValueChanged(nil, logs, nil, nil)
	if err != nil {
		log.Fatal("WatchValueChanged: ", err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-logs:
			fmt.Printf("Sender %s set value to %s\n", log.Sender, log.Value)
		}
	}
}
