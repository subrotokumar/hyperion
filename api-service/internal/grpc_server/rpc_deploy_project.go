package grpc_server

import (
	"context"
	"hyperion/internal/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) DeployProject(ctx context.Context, req *pb.DeployProjectRequest) (*pb.DeployProjectResponse, error) {
	if false {
		return nil, status.Errorf(codes.Internal, "Failed to hash password")
	}
	resp := &pb.DeployProjectResponse{
		Id:        34,
		CreatedBy: req.GetCreatedBy(),
		Name:      req.GetName(),
		GithubUrl: req.GetCreatedBy(),
		Subdomain: req.GetCreatedBy(),
		CreatedAt: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	}
	return resp, nil
}
