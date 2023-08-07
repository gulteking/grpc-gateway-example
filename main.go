package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/gulteking/grpc-gateway-example/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedExampleServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Hello(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	message := fmt.Sprintf("Hello %s", in.Name)
	if in.Email != nil {
		message = fmt.Sprintf("%s, your email is %s", message, *in.Email)
	}
	return &pb.HelloResponse{
		Message: message,
	}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			validator.UnaryServerInterceptor(),
		))
	// Attach the Greeter service to the server
	pb.RegisterExampleServer(s, NewServer())
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8087")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8087",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterExampleHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8088",
		Handler: gwmux,
	}
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8088")
	log.Fatalln(gwServer.ListenAndServe())
}
