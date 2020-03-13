package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-blockchain/proto"
	"grpc-blockchain/server/blockchain"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen port 8080: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockChainServer(srv, &Server{
		BlockChain: blockchain.NewBlockChain(),
	})
	_ = srv.Serve(listener)
}

// Server implements proto.BlockChainServer interface
type Server struct {
	BlockChain *blockchain.BlockChain
}

// AddBlock : adds new block to blockchain
func (s *Server) AddBlock(_ context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.BlockChain.AddBlock(in.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockChain : returns blockChain
func (s *Server) GetBlockChain(_ context.Context, _ *proto.GetBlockChainRequest) (*proto.GetBlockChainResponse, error) {
	resp := new(proto.GetBlockChainResponse)
	for _, b := range s.BlockChain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Data:          b.Data,
			Hash:          b.Hash,
		})
	}

	return resp, nil
}
