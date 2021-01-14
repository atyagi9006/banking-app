module github.com/atyagi9006/banking-app/account-svc

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef
	github.com/atyagi9006/banking-app/auth-mgr-svc v0.0.0-20210114172434-76167971f872
	github.com/atyagi9006/opa-authz v0.0.0-20210102190518-cd21445d876e
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/lib/pq v1.9.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	google.golang.org/genproto v0.0.0-20210108203827-ffc7fda8c3d7
	google.golang.org/grpc v1.34.0
)
