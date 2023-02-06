package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"tiny-tiktok/service/video/internal/config"
	"tiny-tiktok/service/video/internal/handler"
	"tiny-tiktok/service/video/internal/svc"
)

var configFile = flag.String("f", "etc/video-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	//model.NewDB()
	//model.DB.AutoMigrate(&Video{})
}
