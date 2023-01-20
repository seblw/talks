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
		nodeAddr     = "http://127.0.0.1:8545/"
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
	val, err := instance.Value(nil)
	if err != nil {
		log.Fatal("SetValue: ", err)
	}
	fmt.Println("Value: ", val)
}
