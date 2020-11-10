package server

//import (
//	"chick/api/oauth2"
//	"chick/micro-svr/oauth2/conf"
//	"chick/micro-svr/oauth2/service"
//	"chick/pkg/log"
//	"chick/pkg/micro"
//
//	"github.com/micro/go-micro/v2/server"
//)
//
//func Init(c *conf.Config, svr *service.Service) {
//	micro.InitServer(c.Name, "latest", c.Registry, nil, func(s server.Server) {
//		if err := oauth2.RegisterOauth2Handler(s, svr); err != nil {
//			log.Panic(err)
//		}
//	})
//}
