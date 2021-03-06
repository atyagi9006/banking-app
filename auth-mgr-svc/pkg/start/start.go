package start

import (
	"fmt"
	"log"
	"net"

	"github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/api"
	pb "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addressFlag  = flag.String("acc-grpc-ip", "", "gRPC listening IP")
	portFlag     = flag.Uint16("grpc-port", 7781, "gRPC listening port")
	gwPortFlag   = flag.Uint16("rest-port", 7782, "REST gateway port")
	certFile     = "pkg/cert/server.crt"
	keyFile      = "pkg/cert/server.key"
	insecureFlag = flag.Bool("insecure", true, "Run in insecure mode")
)

func Run() {

	grpcAddress := grpcAddressStr()
	restAddress := restAddressStr()

	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpcAddress)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	/* go func() {
		err := startRESTServer(grpcAddress, restAddress)
		if err != nil {
			log.Fatalf("failed to start gRPC  GW server: %s", err)
		}
	}() */

	log.Println("GRPC-server started at ", grpcAddress)
	log.Println("GRPC-GW-server started at ", restAddress)
	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}

func startGRPCServer(grpcAddress string) error {
	//create a listener on tcp layer
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// create service hello service
	authSvc, err := api.NewAuthMgrService()
	if err != nil {
		log.Fatal(err)
	}

	// Create an array of gRPC options with the credentials
	grpcOpts := setupGrpcServerOptions() //account service is used as a unaryInterceptor

	//create grpc service
	grpcServer := grpc.NewServer(grpcOpts...)

	pb.RegisterAuthMgrServiceServer(grpcServer, authSvc)

	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", grpcAddress)
	//start grpc server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}
	return nil
}

func setupGrpcServerOptions() []grpc.ServerOption {
	if !*insecureFlag {
		// Create the TLS credentials
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("could not load TLS keys: %s", err)
		}
		return []grpc.ServerOption{grpc.Creds(creds),
			grpc.UnaryInterceptor(unaryInterceptor)}
	}
	// This is where you can setup custom options for the grpc server
	// https://godoc.org/google.golang.org/grpc#ServerOption

	return []grpc.ServerOption{grpc.UnaryInterceptor(unaryInterceptor)}
}

func setupServeMuxOptions() []runtime.ServeMuxOption {
	return nil
}

func setupGrpcDialOptions() []grpc.DialOption {
	if !*insecureFlag {
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			fmt.Errorf("could not load TLS certificate: %s", err)
		}
		// Setup the client gRPC options
		return []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	}
	// This is where you can set up your dial options.
	// https://godoc.org/google.golang.org/grpc#DialOption
	return []grpc.DialOption{grpc.WithInsecure()}
}

//gRPCAddress
func grpcAddressStr() string {
	if *addressFlag == "" {
		*addressFlag = "localhost"
	}
	return fmt.Sprintf("%s:%d", *addressFlag, *portFlag)
}

func restAddressStr() string {
	if *addressFlag == "" {
		*addressFlag = "localhost"
	}
	return fmt.Sprintf("%s:%d", *addressFlag, *gwPortFlag)
}
