package composer

import (
	"demo-service/common"
	"demo-service/proto/pb"
	authBusiness "demo-service/services/auth/business"
	authSQLStorage "demo-service/services/auth/storage/mysql"
	authUserRPC "demo-service/services/auth/storage/rpc"
	authAPI "demo-service/services/auth/transport/api"
	authRPC "demo-service/services/auth/transport/rpc"
	taskBusiness "demo-service/services/task/business"
	taskRepository "demo-service/services/task/repository"
	taskSQLStorage "demo-service/services/task/storage/mysql"
	taskUserRPC "demo-service/services/task/storage/rpc"
	taskAPI "demo-service/services/task/transport/api"
	userBusiness "demo-service/services/user/business"
	userSQLStorage "demo-service/services/user/storage/mysql"
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

	userStore := userSQLStorage.NewMySQLStore(db.GetDB())
	// In this case, user business does not need repository layer
	biz := userBusiness.NewBusiness(userStore)
	userService := userApi.NewAPI(biz)

	return userService
}

func ComposeTaskAPIService(serviceCtx sctx.ServiceContext) TaskService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	taskStore := taskSQLStorage.NewMySQLStore(db.GetDB())
	userClient := taskUserRPC.NewClient(composeUserRPCClient(serviceCtx))

	repo := taskRepository.NewRepository(taskStore, userClient)
	biz := taskBusiness.NewBusiness(repo)
	serviceAPI := taskAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeAuthAPIService(serviceCtx sctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authStore := authSQLStorage.NewMySQLStore(db.GetDB())
	hasher := new(common.Hasher)

	userClient := authUserRPC.NewClient(composeUserRPCClient(serviceCtx))

	// In this case, auth business does not need repository layer
	// instead, we can pass auth store because it implements repository interface
	biz := authBusiness.NewBusiness(authStore, userClient, jwtComp, hasher)
	serviceAPI := authAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userStore := userSQLStorage.NewMySQLStore(db.GetDB())
	// In this case, user business does not need repository layer
	userBusiness := userBusiness.NewBusiness(userStore)
	userService := userRPC.NewService(userBusiness)

	return userService
}

func ComposeAuthGRPCService(serviceCtx sctx.ServiceContext) pb.AuthServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authStore := authSQLStorage.NewMySQLStore(db.GetDB())
	hasher := new(common.Hasher)

	// In Auth GRPC service, user repository is unnecessary
	biz := authBusiness.NewBusiness(authStore, nil, jwtComp, hasher)
	authService := authRPC.NewService(biz)

	return authService
}
