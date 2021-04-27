package main

import (
	"log"
	"net"

	"github.com/roaires/FullCycle-gRPC-Golang/pb"
	"github.com/roaires/FullCycle-gRPC-Golang/services"
	"google.golang.org/grpc"
)

func main() {

	// Definição da port para o server
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
	}

	//Criação do server
	grpcServer := grpc.NewServer()

	//Registrar o service
	pb.RegisterUserServiceServer(grpcServer, &services.UserService{})

	//Para possibilitar testes utilizando projeto Evans
	//reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Erro ao conectar ao serve: %v", err)
	}

}
