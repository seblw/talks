package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seblw/eth-gowroc/pkg/storage"
)

func main() {
	var (
		nodeAddr     = "http://127.0.0.1:8545/"
		privKey      = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
		contractAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	)
	valueStr := os.Args[1]
	val, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		log.Fatal("ParseInt: ", err)
	}

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
	addr := common.HexToAddress(contractAddr)
	instance, err := storage.NewStorage(addr, cli)
	if err != nil {
		log.Fatal("NewStorage: ", err)
	}
	tx, err := instance.SetValue(auth, big.NewInt(val))
	if err != nil {
		log.Fatal("SetValue: ", err)
	}
	fmt.Println("Tx: ", tx.Hash())
}
