package user

import (
	"context"
	"demo-service/proto/pb"
	"demo-service/services/user/entity"
	"github.com/viettranx/service-context/core"
)

type Business interface {
	GetUserDetails(ctx context.Context, id int) (*entity.User, error)
	GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error)
	CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (s *grpcService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*pb.PublicUserInfoResp, error) {
	user, err := s.business.GetUserDetails(ctx, int(req.Id))

	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.PublicUserInfoResp{
		User: &pb.PublicUserInfo{
			Id:        int32(user.Id),
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, nil
}

func (s *grpcService) GetUsersByIds(ctx context.Context, req *pb.GetUsersByIdsReq) (*pb.PublicUsersInfoResp, error) {
	userIDs := make([]int, len(req.Ids))

	for i := range userIDs {
		userIDs[i] = int(req.Ids[i])
	}

	users, err := s.business.GetUsersByIds(ctx, userIDs)

	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	publicUserInfo := make([]*pb.PublicUserInfo, len(users))

	for i := range users {
		publicUserInfo[i] = &pb.PublicUserInfo{
			Id:        int32(users[i].Id),
			FirstName: users[i].FirstName,
			LastName:  users[i].LastName,
		}
	}

	return &pb.PublicUsersInfoResp{Users: publicUserInfo}, nil
}

func (s *grpcService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.NewUserIdResp, error) {
	newUserData := entity.NewUserForCreation(req.FirstName, req.LastName, req.Email)

	if err := s.business.CreateNewUser(ctx, &newUserData); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.NewUserIdResp{Id: int32(newUserData.Id)}, nil
}
