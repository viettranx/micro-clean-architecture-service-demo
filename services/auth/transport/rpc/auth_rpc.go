package rpc

import (
	"context"
	"demo-service/proto/pb"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type Business interface {
	IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) IntrospectToken(ctx context.Context, req *pb.IntrospectReq) (*pb.IntrospectResp, error) {
	claims, err := s.business.IntrospectToken(ctx, req.AccessToken)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.IntrospectResp{
		Tid: claims.ID,
		Sub: claims.Subject,
	}, nil
}
