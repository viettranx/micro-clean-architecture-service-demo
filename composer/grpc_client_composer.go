package composer

import (
	"context"
	"demo-service/common"
	"demo-service/proto/pb"
	sctx "github.com/viettranx/service-context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type authClient struct {
	grpcAuthClient pb.AuthServiceClient
}

func (ac *authClient) IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error) {
	resp, err := ac.grpcAuthClient.IntrospectToken(ctx, &pb.IntrospectReq{AccessToken: accessToken})

	if err != nil {
		return "", "", err
	}

	return resp.Sub, resp.Tid, nil
}

// ComposeAuthRPCClient use only for middleware: get token info
func ComposeAuthRPCClient(serviceCtx sctx.ServiceContext) *authClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCServerAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return &authClient{pb.NewAuthServiceClient(clientConn)}
}

func composeUserRPCClient(serviceCtx sctx.ServiceContext) pb.UserServiceClient {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCServerAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return pb.NewUserServiceClient(clientConn)
}
