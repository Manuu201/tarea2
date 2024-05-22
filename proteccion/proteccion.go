package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Manuu201/tarea2/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedClienteServiceServer
}

func (s *server) AccionCliente(ctx context.Context, req *pb.ClienteRequest) (*pb.ClienteResponse, error) {
	// Implementa la lógica de tu servicio
	mensaje := fmt.Sprintf("Recibido: Rut: %s, Correo: %s, PDF size: %d bytes", req.Rut, req.Correo, len(req.Pdf))
	return &pb.ClienteResponse{Mensaje: mensaje}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterClienteServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
