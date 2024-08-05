package main

import (
	"context"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	server "cloudbees_assignmnet.com/ass1/Mapper"
	"cloudbees_assignmnet.com/ass1/internal/App"
	service "cloudbees_assignmnet.com/ass1/internal/App/services"
	"cloudbees_assignmnet.com/ass1/internal/probuf/probuf_generated"
)

func main() {
	myApp := fx.New(
		fx.Provide(
			context.Background,
		),
		App.Module,
		server.Module,
		fx.Invoke(initializeTrains),
		fx.Invoke(startServer),
	)

	ctx := context.Background()

	if err := myApp.Start(ctx); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
	defer myApp.Stop(ctx)

	select {}

}

func initializeTrains(trainService service.TrainService) {
	trainService.InitTrains()
}

func startServer(localServer probuf_generated.RailwayServiceServer) {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	probuf_generated.RegisterRailwayServiceServer(server, localServer)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
