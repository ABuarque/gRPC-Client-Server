package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ABuarque/blockchain/proto"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add a new block")
	listFlag := flag.Bool("list", false, "list blockchain")

	flag.Parse()

	con, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot connect to server")
	}

	client = proto.NewBlockchainClient(con)

	if *addFlag {
		AddBlock()
	}

	if *listFlag {
		ListBlockchain()
	}
}

func AddBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})

	if err != nil {
		log.Fatal("Unable to add block")
	}
	fmt.Printf("Added block with hash %s", block.Hash)
}

func ListBlockchain() {
	blockchain, err := client.GetBlockChain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatal("Unable to get blockchain")
	}
	for _, b := range blockchain.Blocks {
		fmt.Printf("Hash: %s, data: %s\n", b.Hash, b.Data)
	}
}
