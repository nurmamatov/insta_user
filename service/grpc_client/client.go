package grpcclient

import (
	"fmt"
	"tasks/Instagram_clone/insta_user/config"
	pp "tasks/Instagram_clone/insta_user/genproto/post_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	PostService() pp.PostServiceClient
}

// Client
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"post_service":    pp.NewPostServiceClient(connPost),
		},
	}, nil
}

func (g *GrpcClient) PostService() pp.PostServiceClient {
	return g.connections["post_service"].(pp.PostServiceClient)
}
