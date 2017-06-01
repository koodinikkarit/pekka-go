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

	"fmt"

	"github.com/koodinikkarit/pekka/db"
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

func (s *server) CreateWeeklyTimer(ctx context.Context, in *PekkaService.CreateWeeklyTimerRequest) (*PekkaService.CreateWeeklyTimerResponse, error) {
	return &PekkaService.CreateWeeklyTimerResponse{}, nil
}

func (s *server) EditWeeklyTimer(ctx context.Context, in *PekkaService.EditWeeklyTimerRequest) (*PekkaService.EditWeeklyTimerResponse, error) {
	return &PekkaService.EditWeeklyTimerResponse{}, nil
}

func (s *server) FetchExecutors(fetchExecutors *PekkaService.FetchExecutorsRequest, stream PekkaService.Pekka_FetchExecutorsServer) error {
	return nil
}

func (s *server) FetchExecutorById(ctx context.Context, in *PekkaService.FetchExecutorByIdRequest) (*PekkaService.FetchExecutorByIdResponse, error) {
	return &PekkaService.FetchExecutorByIdResponse{}, nil
}

func (s *server) FetchExecutorActionsByExecutorId(in *PekkaService.FetchExecutorActionsByExecutorIdRequest, stream PekkaService.Pekka_FetchExecutorActionsByExecutorIdServer) error {
	return nil
}

func (s *server) CreateExecutor(ctx context.Context, in *PekkaService.CreateExecutorRequest) (*PekkaService.CreateExecutorResponse, error) {
	return &PekkaService.CreateExecutorResponse{}, nil
}

func (s *server) AddExecutorActionToExecutor(ctx context.Context, in *PekkaService.AddExecutorActionToExecutorRequest) (*PekkaService.AddExecutorActionToExecutorResponse, error) {
	return &PekkaService.AddExecutorActionToExecutorResponse{}, nil
}

func (s *server) EditButton(ctx context.Context, in *PekkaService.EditButtonRequest) (*PekkaService.EditButtonResponse, error) {
	return &PekkaService.EditButtonResponse{}, nil
}

func (s *server) FetchPenttiDevices(in *PekkaService.FetchPenttiDevicesRequest, stream PekkaService.Pekka_FetchPenttiDevicesServer) error {
	penttiDevices := []pekka.Pentti{}
	s.db.Find(&penttiDevices)

	for _, pentti := range penttiDevices {
		stream.Send(&PekkaService.Pentti{
			Id: pentti.ID,
			Ip: pentti.Ip,
		})
	}

	return nil
}

func (s *server) FetchPenttiById(ctx context.Context, in *PekkaService.FetchPenttiByIdRequest) (*PekkaService.FetchPenttiByIdResponse, error) {
	var pentti pekka.Pentti
	s.db.First(&pentti, in.PenttiId)
	return &PekkaService.FetchPenttiByIdResponse{
		Pentti: &PekkaService.Pentti{
			Id: pentti.ID,
			Ip: pentti.Ip,
		},
	}, nil
}

func (s *server) FetchButtonsByPenttiId(in *PekkaService.FetchButtonsByPenttiIdRequest, stream PekkaService.Pekka_FetchButtonsByPenttiIdServer) error {
	buttons := []pekka.Button{}
	s.db.Find(&buttons)

	for _, button := range buttons {
		stream.Send(&PekkaService.Button{
			Id:       button.ID,
			PenttiId: button.PenttiID,
			Number:   uint32(button.ButtonNumber),
		})
	}
	return nil
}

func CreateService(db *gorm.DB, port string) *grpc.Server {
	fmt.Println("create grpc")
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
