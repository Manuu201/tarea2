package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Manuu201/tarea2/proto"
	"google.golang.org/grpc"
)

func main() {
	// Establecer la conexión con el servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// Crear un cliente gRPC
	c := pb.NewClienteServiceClient(conn)

	// Definir los datos del cliente
	rut := "12345678-9"
	correo := "ejemplo@correo.com"

	// Llamar al servicio de AccionCliente con los datos del cliente
	response, err := c.AccionCliente(context.Background(), &pb.ClienteRequest{
		Rut:    rut,
		Correo: correo,
		// Omitir el campo Pdf
	})
	if err != nil {
		log.Fatalf("error calling AccionCliente: %v", err)
	}

	// Imprimir la respuesta del servidor
	fmt.Println("Response:", response.Mensaje)
}
