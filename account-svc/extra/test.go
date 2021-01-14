package main

import (
	"context"
	"log"

	authmgrClient "github.com/atyagi9006/banking-app/auth-mgr-svc/client"
	authmgrPB "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
)

func main() {
	authmgrClient, err := authmgrClient.NewClient(":7781")
	if err != nil {
		log.Fatal("cannot connect to account-svc", err)
	}
	//register user to auth
	registerReq := authmgrPB.RegisterUserRequest{
		Email:    "a.a@c.b",
		Password: "test",
		Role:     "role",
	}
	ctx := context.Background()
	regResp, err := authmgrClient.RegisterUser(ctx, &registerReq)
	if err != nil {
		log.Println("Error with create new employee failed while auth reg: ", err)
	}
	log.Println("user is registered successfully on auth ID:", regResp.Id)
}
