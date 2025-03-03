package server

import (
	"context"
	"log"
	"net"

	"github.com/steve-westwood/explore-service/src/explore_service"
	"github.com/steve-westwood/explore-service/src/persistence"
	"google.golang.org/grpc"
)

type Controllers struct {
	explore_service.UnimplementedExploreServiceServer
	recipientService persistence.RecipientService
}

func (s *Controllers) ListLikedYou(ctx context.Context, in *explore_service.ListLikedYouRequest) (*explore_service.ListLikedYouResponse, error) {
	likedYouList, err := s.recipientService.ListLikedYou(in.RecipientUserId)
	if err != nil {
		return nil, err
	}
	likers := make([]*explore_service.ListLikedYouResponse_Liker, 0, len(*likedYouList.Likers))
	for _, liker := range *likedYouList.Likers {
		likers = append(likers, &explore_service.ListLikedYouResponse_Liker{
			ActorId:       liker.ActorID,
			UnixTimestamp: liker.TimeStamp,
		})
	}
	return &explore_service.ListLikedYouResponse{
		Likers: likers,
	}, nil
}

func (s *Controllers) ListNewLikedYou(ctx context.Context, in *explore_service.ListLikedYouRequest) (*explore_service.ListLikedYouResponse, error) {
	likedYouList, err := s.recipientService.ListNewLikedYou(in.RecipientUserId)
	if err != nil {
		return nil, err
	}
	likers := make([]*explore_service.ListLikedYouResponse_Liker, 0, len(*likedYouList.Likers))
	for _, liker := range *likedYouList.Likers {
		likers = append(likers, &explore_service.ListLikedYouResponse_Liker{
			ActorId:       liker.ActorID,
			UnixTimestamp: liker.TimeStamp,
		})
	}
	return &explore_service.ListLikedYouResponse{
		Likers: likers,
	}, nil
}

func (s *Controllers) CountLikedYou(ctx context.Context, in *explore_service.CountLikedYouRequest) (*explore_service.CountLikedYouResponse, error) {
	count, err := s.recipientService.CountLikedYou(in.RecipientUserId)
	if err != nil {
		return nil, err
	}
	return &explore_service.CountLikedYouResponse{
		Count: uint64(count),
	}, nil
}

func (s *Controllers) PutDecision(ctx context.Context, in *explore_service.PutDecisionRequest) (*explore_service.PutDecisionResponse, error) {
	reciprocated, err := s.recipientService.PutDecision(in.RecipientUserId, in.ActorUserId, in.LikedRecipient)
	if err != nil {
		return nil, err
	}
	return &explore_service.PutDecisionResponse{
		MutualLikes: reciprocated,
	}, nil
}

func NewServer(recipientService persistence.RecipientService) *grpc.Server {
	serverRegistrar := grpc.NewServer()
	service := &Controllers{recipientService: recipientService}
	explore_service.RegisterExploreServiceServer(serverRegistrar, service)
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
