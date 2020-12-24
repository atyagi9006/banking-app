# Banking App 

Banking app is a simple solution for running a  bank with minimal features.

# Repository Structure

## **account-svc**

This service is heart of banking system at current stage. It provides api to manage employees of bank, manage customers and theirs accounts.
It also manges(authorization) access to apis as well as authentication via bearer token (JWT).

## **auth-svc**
This service provides sign in and sign out apis. 

# How to run Banking App

### **pre requisites**
1. Install docker and run :
    ```
    docker pull postgres:12-alpine 
    ```
1. postgres : postgres:12-alpine(docker image)
1. Install golang and configure $GOPATH, $GOROOT
1. Run make init (this will install go-migrate)
1. Run for initial setup : 
    ```
    docker run --name dev-postgres -p 5432:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=P@ssw0rd -d postgres:12-alpine
    ```

### **Steps**
1. ```cd account-svc``` and run : 
    ```
    make startdb
    make build-server
    make run
    ```
1. ```cd auth-svc``` and run :
    ```
    make build-server
    make run
    ```
1. ```Import both postman collection and environment file in postman and run any request. ``` 


## Project Tree
``` 
├── README.md
├── account-svc
│   ├── Makefile
│   ├── Makefile.bank
│   ├── auth
│   │   └── jwt_manager.go
│   ├── bin
│   │   ├── client
│   │   ├── client_main
│   │   └── server
│   ├── client
│   │   └── client.go
│   ├── db
│   │   ├── bank_employee.sql.go
│   │   ├── customer.sql.go
│   │   ├── customer_accounts.sql.go
│   │   ├── db.go
│   │   ├── entry.sql.go
│   │   ├── models.go
│   │   ├── store.go
│   │   ├── store_test.go
│   │   ├── test_util.go
│   │   └── transfer.sql.go
│   ├── go.mod
│   ├── go.sum
│   ├── pkg
│   │   ├── api
│   │   │   ├── accoutnservice.go
│   │   │   ├── customer_account_handler.go
│   │   │   ├── customer_account_handler_test.go
│   │   │   ├── customer_handles.go
│   │   │   ├── customer_hundler_test.go
│   │   │   ├── handler.go
│   │   │   ├── handler_test.go
│   │   │   ├── intercptor.go
│   │   │   └── token_mgr_handler.go
│   │   ├── cert
│   │   │   ├── server.crt
│   │   │   ├── server.csr
│   │   │   └── server.key
│   │   ├── proto
│   │   │   ├── api.pb.go
│   │   │   ├── api.pb.gw.go
│   │   │   ├── api.proto
│   │   │   └── api.swagger.json
│   │   └── start
│   │       ├── auth.go
│   │       └── start.go
│   ├── scripts
│   │   └── db
│   │       ├── migration
│   │       │   ├── 000001_init_bank_emp_schema.down.sql
│   │       │   ├── 000001_init_bank_emp_schema.up.sql
│   │       │   ├── 000002_init_bank_schema.down.sql
│   │       │   └── 000002_init_bank_schema.up.sql
│   │       └── query
│   │           ├── bank_employee.sql
│   │           ├── customer.sql
│   │           ├── customer_accounts.sql
│   │           ├── entry.sql
│   │           └── transfer.sql
│   ├── server
│   │   └── server.go
│   └── sqlc.yaml
├── auth-svc
│   ├── Makefile
│   ├── bin
│   │   └── server
│   ├── client
│   │   └── client.go
│   ├── go.mod
│   ├── go.sum
│   ├── pkg
│   │   ├── api
│   │   │   └── handler.go
│   │   ├── cert
│   │   │   ├── server.crt
│   │   │   ├── server.csr
│   │   │   └── server.key
│   │   ├── proto
│   │   │   ├── api.pb.go
│   │   │   ├── api.pb.gw.go
│   │   │   ├── api.proto
│   │   │   └── api.swagger.json
│   │   └── start
│   │       ├── auth.go
│   │       └── start.go
│   └── server
│       └── server.go
└── postman
    ├── bank-local.postman_environment.json
    └── my_bank.postman_collection.json 
```


