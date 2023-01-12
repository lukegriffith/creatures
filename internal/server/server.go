package server

import (
	"context"
	"errors"

	pb "github.com/lukegriffith/creatures/pkg/proto"
)

type WorldServer struct {
	pb.UnimplementedWorldServiceServer
}

func NewServer() *WorldServer {
	return nil
}

func (s *WorldServer) CreateWorld(ctx context.Context, w *pb.World) (*pb.WorldResponse, error) {
	return nil, errors.New("Not implemented")
}

// Sets a singletons value to the selected world
func (s *WorldServer) SelectWorld(ctx context.Context, w *pb.WorldSelectionRequest) (*pb.WorldResponse, error) {
	return nil, errors.New("Not implemented")
}

// Returns a list of all created worlds
// If an ID is provided, it returns a single world
func (s *WorldServer) GetWorld(ctx context.Context, w *pb.WorldRequest) (*pb.WorldResponse, error) {
	return nil, errors.New("Not implemented")
}

// Sets a singletons value to the selected world

// Returns a list of all created worlds
// If an ID is provided, it returns a single world
