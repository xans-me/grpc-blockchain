package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-blockchain/proto"
	"log"
	"time"
)

var client proto.BlockChainClient

func main() {
	addFlag := flag.Bool("add", false, "Add new block")
	listFlag := flag.Bool("list", false, "List all blocks")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = proto.NewBlockChainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockChain()
	}
}

func addBlock() {
	block, addErr := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if addErr != nil {
		log.Fatalf("unable to add block: %v", addErr)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockChain() {
	blockChain, getErr := client.GetBlockChain(context.Background(), &proto.GetBlockChainRequest{})
	if getErr != nil {
		log.Fatalf("unable to get blockchain: %v", getErr)
	}

	log.Println("blocks:")
	for _, b := range blockChain.Blocks {
		log.Printf("hash %s, prev hash: %s, data: %s\n", b.Hash, b.PrevBlockHash, b.Data)
	}
}
