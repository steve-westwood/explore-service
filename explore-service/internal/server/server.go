package server

import (
	"log"
	"net"

	"github.com/steve-westwood/explore-service/explore-service/internal/app"
	pb "github.com/steve-westwood/explore-service/shared/proto"
	"google.golang.org/grpc"
)

type Controllers struct {
	pb.UnimplementedExploreServiceServer
	recipientService app.RecipientService
}

func NewServer(recipientService app.RecipientService) *grpc.Server {
	serverRegistrar := grpc.NewServer()
	service := &Controllers{recipientService: recipientService}
	pb.RegisterExploreServiceServer(serverRegistrar, service)
	return serverRegistrar
}

func StartServer(server *grpc.Server) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// ... rest of your server methods ...
