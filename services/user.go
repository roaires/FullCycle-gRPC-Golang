package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/roaires/FullCycle-gRPC-Golang/pb"
)

func newUserService() *UserService {
	return &UserService{}
}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

// Client/Server
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	/* Simulando um insert no DB, por isso atribuição do id manualmente */

	return &pb.User{
		Id:    "1234567",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

// Server Stream
func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	/*Envio do usuário de forma particioinada para simular uso do Stream*/

	stream.Send(&pb.UserResultStream{
		Status: "Instância de um novo objeto User",
		User: &pb.User{
			Id:    "null",
			Name:  "null",
			Email: "null",
		},
	})
	time.Sleep(time.Second * 5)

	stream.Send(&pb.UserResultStream{
		Status: "Objeto user com valores preenchidos",
		User: &pb.User{
			Id:    req.Id,
			Name:  req.Name,
			Email: req.Email,
		},
	})
	time.Sleep(time.Second * 5)

	stream.Send(&pb.UserResultStream{
		Status: "Simulando registro inserido",
		User: &pb.User{
			Id:    "1234",
			Name:  req.Name,
			Email: req.Email,
		},
	})
	time.Sleep(time.Second * 5)

	stream.Send(&pb.UserResultStream{

		Status: "Simulando retorno do registro inserido",
		User: &pb.User{
			Id:    "1234",
			Name:  req.Name,
			Email: req.Email,
		},
	})
	time.Sleep(time.Second * 5)

	return nil
}

// Client Strem
func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {

	fmt.Println("---------------------------- Server: Client Stream -----------------------------")

	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Erro ao receber stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Inserindo", req.GetName())

	}
}

// Stream bi-direcional
func (*UserService) AddUsersStreamBoth(stream pb.UserService_AddUsersStreamBothServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Erro ao receber stream do client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})
		if err != nil {
			log.Fatalf("Erro ao enviar stream do client: %v", err)
		}
	}
}
