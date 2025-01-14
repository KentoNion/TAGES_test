package main

import (
	"google.golang.org/grpc"
	"net"
	server "tages.local/app/gates/server/grpc"
	"tages.local/app/internal/config"
	"tages.local/app/internal/logger"
	"tages.local/grpc/images/pb"
)

func main() {
	//инициализация конфига
	cfg := config.MustLoad()

	//инициализация логгера
	log := logger.MustInitLogger(cfg)
	log.Info("Logger initialized")
	log.Debug("Debug logging enabled")

	//запуск gRPC сервера
	lis, err := net.Listen("tcp", cfg.Grpc.Port)
	if err != nil {
		panic(err)
	}
	log.Info("Listening on port %s", cfg.Grpc.Port)
	grpcServer := grpc.NewServer() // Создаём один экземпляр gRPC-сервера

	pb.RegisterImageServiceServer(grpcServer, server.NewServer(cfg, log)) // Регистрируем сервис

	log.Info("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil { // Используем тот же сервер для запуска
		panic(err)
	}
}
