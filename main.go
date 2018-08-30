package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var transactionHash string = "0x487760edc4640d45a2829b0aa0d41d671726436b5ac73047aa634e1f4f905457"

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Whoops something went wrong!", err)
	}

	ctx := context.Background()
	tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash(transactionHash))
	if !pending {
		spew.Dump(tx)
	}
}
