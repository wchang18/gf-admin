package cmd

import (
	"context"
	"gf-admin/internal/controller/kube"
	"gf-admin/internal/controller/role"
	"gf-admin/internal/controller/user"
	_ "gf-admin/internal/logic"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareCORS)
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				s.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1(),
					)
				})
				group.Bind(role.NewV1())
				s.Group("/kube", func(group *ghttp.RouterGroup) {
					group.Bind(
						kube.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
