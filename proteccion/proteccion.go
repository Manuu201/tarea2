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
	// Registro de información para verificar la solicitud recibida
	log.Printf("Received request: Rut: %s, Correo: %s", req.Rut, req.Correo)

	// Implementa la lógica de tu servicio
	mensaje := fmt.Sprintf("Recibido: Rut: %s, Correo: %s", req.Rut, req.Correo)
	return &pb.ClienteResponse{Mensaje: mensaje}, nil
}

func main() {
	// Configurar el servidor gRPC
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterClienteServiceServer(s, &server{})

	// Iniciar el servidor gRPC
	log.Println("Server started, listening on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
