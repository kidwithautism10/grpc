package auth

import ssov1 "github.com/kidwithautism10/grpcContract/gen/go/sso"

type ServerAPI struct {
	ssov1.UnimplementedAuthServer
}
