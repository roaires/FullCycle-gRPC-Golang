package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/roaires/FullCycle-gRPC-Golang/pb"
	"google.golang.org/grpc"
)

func main() {

	//WithInsecure - Não utilizar em produção
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Não foi possível conectar ao gRPC Server: %v", err)
	}

	//Encerra conexão quando não estiver mais em uso
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	fmt.Println("---------------------------- Client/Server -----------------------------")
	AddUser(client)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("---------------------------- Server Stream -----------------------------")
	AddUserVerbose(client)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("---------------------------- Client Stream -----------------------------")
	AddUsers(client)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("---------------------------- Stream bi-direcional ----------------------")
	AddUsersStreamBoth(client)
}

// Client/Server
func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Rodrigo Aires",
		Email: "roaires@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Não foi possível executar a request gRPC: %v", err)
	}

	fmt.Println(res)
}

// Stream Server
func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Rodrigo Aires",
		Email: "roaires@gmail.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Não foi possível executar a request gRPC: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			// Sem dados para consumo
			break
		}
		if err != nil {
			log.Fatalf("Não foi possível receber a mensagem via stream: %v", err)
		}

		//Apresenta status de cada partionamento processado
		fmt.Println("Status:", stream.Status)
		fmt.Println(stream.GetUser())
		fmt.Println("")
		fmt.Println("")

	}

}

// Stream Client
func AddUsers(client pb.UserServiceClient) {

	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Usuário 1",
			Email: "user1@user.com",
		},
		{
			Id:    "2",
			Name:  "Usuário 2",
			Email: "user2@user.com",
		},
		{
			Id:    "3",
			Name:  "Usuário 3",
			Email: "user3@user.com",
		},
		{
			Id:    "4",
			Name:  "Usuário 4",
			Email: "user4@user.com",
		},
		{
			Id:    "5",
			Name:  "Usuário 5",
			Email: "user5@user.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Erro na criação da requisição : %v", err)
	}
	for _, req := range reqs {
		fmt.Println("Enviando... ", req.GetName())
		stream.Send(req)
		time.Sleep(time.Second * 5)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Erro ao receber a resposta : %v", err)
	}

	fmt.Println("")
	fmt.Println("Lista de usuários - Retorno do server:")
	fmt.Println(res)
}

// Stream bi-direcional
func AddUsersStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUsersStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Erro na criação da requisição : %v", err)
	}

	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Usuário 1",
			Email: "user1@user.com",
		},
		{
			Id:    "2",
			Name:  "Usuário 2",
			Email: "user2@user.com",
		},
		{
			Id:    "3",
			Name:  "Usuário 3",
			Email: "user3@user.com",
		},
		{
			Id:    "4",
			Name:  "Usuário 4",
			Email: "user4@user.com",
		},
		{
			Id:    "5",
			Name:  "Usuário 5",
			Email: "user5@user.com",
		},
	}

	//Criação de um chanel para evitar que encerre enquanto estiver recebendo stream
	wait := make(chan int)

	/* Funções anônimas assíncronas para processar o envio e recebimento via Stream  */

	go func() {
		for _, req := range reqs {
			fmt.Println("Enviando... ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Erro ao receber stream do server : %v", err)
				break
			}
			fmt.Printf("Recebendo user %v com status: %v\n ", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
