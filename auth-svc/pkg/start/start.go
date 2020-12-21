package start

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/atyagi9006/banking-app/auth-svc/pkg/api"
	pb "github.com/atyagi9006/banking-app/auth-svc/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Run() {

	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
	certFile := "pkg/cert/server.crt"
	keyFile := "pkg/cert/server.key"
	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpcAddress, certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(restAddress, grpcAddress, certFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	log.Println("GRPC-server started at ", grpcAddress)
	log.Println("GRPC-GW-server started at ", restAddress)
	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}

func startGRPCServer(address, certFile, keyFile string) error {
	//create a listener on tcp layer
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// create service hello service
	helloSvc := api.AuthService{}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
		return err
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds),
		grpc.UnaryInterceptor(unaryInterceptor)}

	//create grpc service
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterAuthServiceServer(grpcServer, &helloSvc)

	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	//start grpc server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}
	return nil
}

func startRESTServer(address, grpcAddress, certFile string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return fmt.Errorf("could not load TLS certificate: %s", err)
	}
	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher))

	err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}
