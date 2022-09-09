// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/xdorro/golang-grpc-base-project/internal/interceptor"
	"github.com/xdorro/golang-grpc-base-project/internal/module/auth/biz"
	"github.com/xdorro/golang-grpc-base-project/internal/module/ping/biz"
	"github.com/xdorro/golang-grpc-base-project/internal/module/user/biz"
	"github.com/xdorro/golang-grpc-base-project/internal/server"
	"github.com/xdorro/golang-grpc-base-project/internal/service"
	"net/http"
)

// Injectors from wire.go:

func initServer() server.IServer {
	serveMux := http.NewServeMux()
	iInterceptor := interceptor.NewInterceptor()
	iPingService := pingbiz.NewService()
	option := &userbiz.Option{}
	iUserService := userbiz.NewService(option)
	authbizOption := &authbiz.Option{}
	iAuthService := authbiz.NewService(authbizOption)
	serviceOption := &service.Option{
		Mux:         serveMux,
		Interceptor: iInterceptor,
		PingService: iPingService,
		UserService: iUserService,
		AuthService: iAuthService,
	}
	iService := service.NewService(serviceOption)
	serverOption := &server.Option{
		Mux:     serveMux,
		Service: iService,
	}
	iServer := server.NewServer(serverOption)
	return iServer
}