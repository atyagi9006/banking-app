package start

import (
	"context"

	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/atyagi9006/banking-app/account-svc/pkg/api"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var (
	certFile      = "pkg/cert/server.crt"
	keyFile       = "pkg/cert/server.key"
	insecureFlag  = flag.Bool("insecure", true, "Run in insecure mode")
	noAuthFlag    = flag.Bool("no-auth", true, "Run with no auth mode")
	addressFlag   = flag.String("acc-grpc-ip", "", "gRPC listening IP")
	portFlag      = flag.Uint16("grpc-port", 7777, "gRPC listening port")
	gwPortFlag    = flag.Uint16("rest-port", 7778, "REST gateway port")
	adminEmail    = "a.tyagi@xyz.com"
	adminPassword = "a.ty@123"
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
	go func() {
		err := startRESTServer(grpcAddress, restAddress)
		if err != nil {
			log.Fatalf("failed to start gRPC  GW server: %s", err)
		}
	}()

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
	accountSvc, err := api.NewAccountService()
	if err != nil {
		return err
	}

	// Create an array of gRPC options with the credentials
	grpcOpts := setupGrpcServerOptions(accountSvc) //account service is used as a unaryInterceptor

	//create grpc service
	grpcServer := grpc.NewServer(grpcOpts...)

	pb.RegisterAccountServiceServer(grpcServer, accountSvc)
	//seed admin
	seedAdmin(accountSvc)
	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", grpcAddress)
	//start grpc server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}
	return nil
}

func startRESTServer(grpcAddress, restAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	svrMuxOpts := setupServeMuxOptions()
	mux := runtime.NewServeMux(svrMuxOpts...)

	dialOpts := setupGrpcDialOptions()
	err := pb.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, grpcAddress, dialOpts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", restAddress)
	http.ListenAndServe(restAddress, mux)
	return nil
}

func setupGrpcServerOptions(interceptor *api.AccountService) []grpc.ServerOption {
	if !*insecureFlag {
		// Create the TLS credentials
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("could not load TLS keys: %s", err)
		}
		return []grpc.ServerOption{grpc.Creds(creds),
			grpc.UnaryInterceptor(interceptor.Unary())}
	}
	// This is where you can setup custom options for the grpc server
	// https://godoc.org/google.golang.org/grpc#ServerOption
	//return nil
	if *noAuthFlag {
		return []grpc.ServerOption{grpc.UnaryInterceptor(unaryInterceptor)}
	}
	return []grpc.ServerOption{grpc.UnaryInterceptor(interceptor.Unary())}
}

func setupServeMuxOptions() []runtime.ServeMuxOption {
	if !*insecureFlag {
		return []runtime.ServeMuxOption{
			runtime.WithIncomingHeaderMatcher(credMatcher),
		}
	}
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

func seedAdmin(svc *api.AccountService) {
	log.Println("Seeding admin...")
	createAdminReq := pb.CreateEmployeeRequest{
		Email:    adminEmail,
		Password: adminPassword,
		FullName: "Admin",
		Role:     "admin",
	}
	_, err := svc.CreateBankEmployee(context.Background(), &createAdminReq)
	if err != nil {
		if codes.AlreadyExists != status.Convert(err).Code() {
			log.Println(err)
		}
	}
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
