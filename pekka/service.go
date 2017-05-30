package pekka

import (
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/pekka/pekka_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type server struct {
	db *gorm.DB
}

func (s *server) FetchWeeklyTimers(fetchWeeklyTimers *PekkaService.FetchWeeklyTimersRequest, stream PekkaService.Pekka_FetchWeeklyTimersServer) error {
	return nil
}

func (s *server) FetchWeeklyTimerById(ctx context.Context, in *PekkaService.FetchWeeklyTimerByIdRequest) (*PekkaService.FetchWeeklyTimerByIdResponse, error) {
	return &PekkaService.FetchWeeklyTimerByIdResponse{}, nil
}

func (s *server) FetchWeeklyTimerBySlug(ctx context.Context, in *PekkaService.FetchWeeklyTimerBySlugRequest) (*PekkaService.FetchWeeklyTimerBySlugResponse, error) {
	return &PekkaService.FetchWeeklyTimerBySlugResponse{}, nil
}

func CreateService(db *gorm.DB, port string) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./ssl/server.crt", "./ssl/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	PekkaService.RegisterPekkaServer(s, &server{db: db})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return s
}
