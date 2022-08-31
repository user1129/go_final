package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Model data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	//write the Log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Model.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	//return response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

func (app *Config) nagRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for GRPC: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Model: app.Models})

	log.Printf("gRPC server started on porf: %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for GRPC: %v", err)
	}
}
