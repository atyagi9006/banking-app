package main

import (
	"context"
	"log"

	"github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//Authentication holds login and password
type Authentication struct {
	Login    string
	Password string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("pkg/cert/server.crt", "localhost")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// Setup the login/pass
	auth := Authentication{
		Login:    "amit",
		Password: "tyagi",
	}

	//get a connection by dialing G-RPC
	conn, err := grpc.Dial(":7777", grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := proto.NewPingClient(conn)
	res, err := client.SayHello(context.Background(), &proto.PingMessage{Greeting: "le greeting ...."})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server : %s \n", res.Greeting)
	res1, err := client.SayHellogw(context.Background(), &proto.PingMessage{Greeting: "le greeting gw ...."})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response1 from server : %s \n", res1.Greeting)

}
