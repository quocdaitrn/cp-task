package adapters

import (
	"context"

	"github.com/quocdaitrn/golang-kit/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/quocdaitrn/cp-task/infra/config"
	"github.com/quocdaitrn/cp-task/proto/pb"
)

type authClient struct {
	grpcAuthClient pb.AuthServiceClient
}

func (ac *authClient) IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error) {
	resp, err := ac.grpcAuthClient.IntrospectToken(ctx, &pb.IntrospectRequest{AccessToken: accessToken})

	if err != nil {
		return "", "", err
	}

	return resp.Sub, resp.Tid, nil
}

// ProvideGRPCAuthClient use only for middleware: get token info
func ProvideGRPCAuthClient(cfg config.Config) (auth.AuthenticateClient, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(cfg.GRPCServerAuthServiceAddress, opts)
	if err != nil {
		return nil, err
	}

	return &authClient{pb.NewAuthServiceClient(clientConn)}, nil
}

func ProvideGRPCUserServiceClient(cfg config.Config) (pb.UserServiceClient, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(cfg.GRPCServerUserServiceAddress, opts)
	if err != nil {
		return nil, err
	}

	return pb.NewUserServiceClient(clientConn), nil
}
