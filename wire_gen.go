// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/cwd-nial/go-statik/config"
	"github.com/cwd-nial/go-statik/driver"
)

import (
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func Server() *driver.Server {
	handler := driver.NewHandler()
	router := driver.NewRouter(handler)
	serverConfig := config.ProvideServerConfig()
	server := driver.NewServer(router, serverConfig)
	return server
}