package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/Manuu201/tarea2/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewClienteServiceClient(conn)

	rut := "12345678-9"
	correo := "ejemplo@correo.com"
	pdfPath := "cliente.pdf"

	pdfData, err := os.ReadFile(pdfPath)
	if err != nil {
		log.Fatalf("error reading PDF: %v", err)
	}

	response, err := c.AccionCliente(context.Background(), &pb.ClienteRequest{
		Rut:    rut,
		Correo: correo,
		Pdf:    pdfData,
	})
	if err != nil {
		log.Fatalf("error calling AccionCliente: %v", err)
	}

	fmt.Println("Response:", response.Mensaje)
}
