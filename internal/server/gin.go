package server

import (
	gin2 "github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	v1 "test/api/helloworld/v1"
	"test/internal/conf"
	"test/internal/service"
	"test/third_party/gin"
)

func NewGinServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *gin.Server {
	srv := gin.NewServer(
		gin.WithAddress(c.Grpc.Addr),
	)
	srv.GET("/helloworld/:name", func(ctx *gin2.Context) {
		var in v1.HelloRequest
		in.Name = ctx.Param("name")
		reply, err := greeter.SayHello(ctx, &in)
		if err != nil {
			ctx.String(200, "err")
			return
		}
		ctx.String(200, reply.Message)
	})
	return srv
}
