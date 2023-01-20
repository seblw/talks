package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seblw/eth-gowroc/pkg/storage"
)

func main() {
	var (
		nodeAddr = "http://127.0.0.1:8545/"
		privKey  = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	)

	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, nodeAddr)
	if err != nil {
		log.Fatal("DialContext: ", err)
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal("HexToECDSA: ", err)
	}

	chainID, err := cli.ChainID(ctx)
	if err != nil {
		log.Fatal("ChainID: ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("NewKeyedTransactorWithChainID: ", err)
	}
	addr, tx, instance, err := storage.DeployStorage(auth, cli)
	if err != nil {
		log.Fatal("DeployStorage: ", err)
	}
	fmt.Println("Contract address: ", addr)
	fmt.Println("Deploy Tx: ", tx.Hash())
	_ = instance
}
