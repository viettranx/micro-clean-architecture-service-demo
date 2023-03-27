package composer

import (
	"demo-service/common"
	"demo-service/proto/pb"
	authBusiness "demo-service/services/auth/business"
	authSQLRepository "demo-service/services/auth/repository/mysql"
	authUserRPC "demo-service/services/auth/repository/rpc"
	authAPI "demo-service/services/auth/transport/api"
	authRPC "demo-service/services/auth/transport/rpc"
	taskBusiness "demo-service/services/task/business"
	taskSQLRepository "demo-service/services/task/repository/mysql"
	taskUserRPC "demo-service/services/task/repository/rpc"
	taskAPI "demo-service/services/task/transport/api"
	userBusiness "demo-service/services/user/business"
	userSQLRepository "demo-service/services/user/repository/mysql"
	userApi "demo-service/services/user/transport/api"
	userRPC "demo-service/services/user/transport/rpc"
	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
)

type TaskService interface {
	CreateTaskHdl() func(*gin.Context)
	GetTaskHdl() func(*gin.Context)
	ListTaskHdl() func(*gin.Context)
	UpdateTaskHdl() func(*gin.Context)
	DeleteTaskHdl() func(*gin.Context)
}

type UserService interface {
	GetUserProfileHdl() func(*gin.Context)
}

type AuthService interface {
	LoginHdl() func(*gin.Context)
	RegisterHdl() func(*gin.Context)
}

func ComposeUserAPIService(serviceCtx sctx.ServiceContext) UserService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())
	biz := userBusiness.NewBusiness(userRepo)
	userService := userApi.NewAPI(biz)

	return userService
}

func ComposeTaskAPIService(serviceCtx sctx.ServiceContext) TaskService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userClient := taskUserRPC.NewClient(composeUserRPCClient(serviceCtx))
	taskRepo := taskSQLRepository.NewMySQLRepository(db.GetDB())
	biz := taskBusiness.NewBusiness(taskRepo, userClient)
	serviceAPI := taskAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeAuthAPIService(serviceCtx sctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authRepo := authSQLRepository.NewMySQLRepository(db.GetDB())
	hasher := new(common.Hasher)

	userClient := authUserRPC.NewClient(composeUserRPCClient(serviceCtx))
	biz := authBusiness.NewBusiness(authRepo, userClient, jwtComp, hasher)
	serviceAPI := authAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())
	userBiz := userBusiness.NewBusiness(userRepo)
	userService := userRPC.NewService(userBiz)

	return userService
}

func ComposeAuthGRPCService(serviceCtx sctx.ServiceContext) pb.AuthServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authRepo := authSQLRepository.NewMySQLRepository(db.GetDB())
	hasher := new(common.Hasher)

	// In Auth GRPC service, user repository is unnecessary
	biz := authBusiness.NewBusiness(authRepo, nil, jwtComp, hasher)
	authService := authRPC.NewService(biz)

	return authService
}
