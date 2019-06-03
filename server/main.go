package main

import (
	"context"
	"log"
	"net"

	"github.com/ABuarque/blockchain/proto"
	"github.com/ABuarque/blockchain/server/blockchain"

	"google.golang.org/grpc"
)

type Server struct {
	Blockchain *blockchain.Blockchain
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Not possible to listen on 8080: %v", err)
	}

	server := grpc.NewServer()

	proto.RegisterBlockchainServer(server, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	server.Serve(listener)
}

func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *Server) GetBlockChain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, block := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: block.PreviousBlockHash,
			Hash:          block.Hash,
			Data:          block.Data,
		})
	}
	return resp, nil
}
