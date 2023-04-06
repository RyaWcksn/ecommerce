package main

import (
	"github.com/RyaWcksn/ecommerce/configs"
	"github.com/RyaWcksn/ecommerce/pkgs/logger"
	"github.com/RyaWcksn/ecommerce/server"
)

func main() {
	cfg := configs.Cfg
	log := logger.New(cfg.App.ENV, cfg.App.APPNAME, cfg.App.LOGLEVEL)
	s := server.NewService(cfg, log)
	s.Start()
}
