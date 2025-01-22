package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"ldm/common/constant"
	"ldm/common/gen_proto/user"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		rsp, err := user.NewUserService(constant.RPC_USER_SRV, micro.NewService().Client()).Profile(context.Background(), &user.ProfileReq{
			Id: 1,
		})
		if err != nil {
			c.JSON(300, err)
			return
		}
		c.JSON(200, rsp.Name)
		return
	})
	r.Run(":8081")
}
