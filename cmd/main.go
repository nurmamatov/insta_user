package main

import (
	"net"
	"tasks/Instagram_clone/insta_user/config"
	pu "tasks/Instagram_clone/insta_user/genproto/user_proto"
	"tasks/Instagram_clone/insta_user/pkg/db"
	"tasks/Instagram_clone/insta_user/pkg/logger"
	"tasks/Instagram_clone/insta_user/service"
	grpcClient "tasks/Instagram_clone/insta_user/service/grpc_client"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.Load()
	log := logger.New(config.LogLevel, "user_service")
	log.Info("main: sqlxConfig",
		logger.String("host", config.PostgresHost),
		logger.Int("port", config.PostgresPort),
		logger.String("database", config.PostgresDatabase))

	grpcClient, err := grpcClient.New(config)
	if err != nil {
		log.Fatal("grpc dial error", logger.Error(err))
	}

	psql, err := db.ConnectToDB(config)
	if err != nil {
		log.Fatal("Error while sqlx connect", logger.Error(err))
	}

	UserService := service.NewPostService(psql, log, grpcClient)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatal("Error while listening:", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pu.RegisterUserServiceServer(s, UserService)
	log.Info("main: server running",
		logger.String("port", config.Port))

	if err = s.Serve(lis); err != nil {
		log.Fatal("Error while listening:", logger.Error(err))
	}
}
