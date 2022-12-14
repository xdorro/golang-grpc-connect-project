// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/xdorro/golang-grpc-base-project/internal/interceptor"
	"github.com/xdorro/golang-grpc-base-project/internal/module/auth/biz"
	"github.com/xdorro/golang-grpc-base-project/internal/module/auth/service"
	"github.com/xdorro/golang-grpc-base-project/internal/module/permission/biz"
	"github.com/xdorro/golang-grpc-base-project/internal/module/permission/service"
	"github.com/xdorro/golang-grpc-base-project/internal/module/role/biz"
	"github.com/xdorro/golang-grpc-base-project/internal/module/role/service"
	"github.com/xdorro/golang-grpc-base-project/internal/module/user/biz"
	"github.com/xdorro/golang-grpc-base-project/internal/module/user/service"
	"github.com/xdorro/golang-grpc-base-project/internal/server"
	"github.com/xdorro/golang-grpc-base-project/internal/service"
	"github.com/xdorro/golang-grpc-base-project/pkg/casbin"
	"github.com/xdorro/golang-grpc-base-project/pkg/redis"
	"github.com/xdorro/golang-grpc-base-project/pkg/repo"
	"net/http"
)

// Injectors from wire.go:

func initServer() server.IServer {
	serveMux := http.NewServeMux()
	iRepo := repo.NewRepo()
	option := &casbin.Option{
		Repo: iRepo,
	}
	iCasbin := casbin.NewCasbin(option)
	iRedis := redis.NewRedis()
	interceptorOption := &interceptor.Option{
		Casbin: iCasbin,
		Redis:  iRedis,
		Repo:   iRepo,
	}
	iInterceptor := interceptor.NewInterceptor(interceptorOption)
	userbizOption := &userbiz.Option{
		Repo: iRepo,
	}
	iUserBiz := userbiz.NewBiz(userbizOption)
	userserviceOption := &userservice.Option{
		UserBiz: iUserBiz,
	}
	iUserService := userservice.NewService(userserviceOption)
	authbizOption := &authbiz.Option{
		Repo: iRepo,
	}
	iAuthBiz := authbiz.NewBiz(authbizOption)
	authserviceOption := &authservice.Option{
		AuthBiz: iAuthBiz,
	}
	iAuthService := authservice.NewService(authserviceOption)
	permissionbizOption := &permissionbiz.Option{
		Repo: iRepo,
	}
	iPermissionBiz := permissionbiz.NewBiz(permissionbizOption)
	permissionserviceOption := &permissionservice.Option{
		PermissionBiz: iPermissionBiz,
	}
	iPermissionService := permissionservice.NewService(permissionserviceOption)
	rolebizOption := &rolebiz.Option{
		Casbin: iCasbin,
	}
	iRoleBiz := rolebiz.NewBiz(rolebizOption)
	roleserviceOption := &roleservice.Option{
		RoleBiz: iRoleBiz,
	}
	iRoleService := roleservice.NewService(roleserviceOption)
	serviceOption := &service.Option{
		Mux:               serveMux,
		Interceptor:       iInterceptor,
		Repo:              iRepo,
		Redis:             iRedis,
		UserService:       iUserService,
		AuthService:       iAuthService,
		PermissionService: iPermissionService,
		RoleService:       iRoleService,
	}
	iService := service.NewService(serviceOption)
	serverOption := &server.Option{
		Mux:     serveMux,
		Service: iService,
	}
	iServer := server.NewServer(serverOption)
	return iServer
}
