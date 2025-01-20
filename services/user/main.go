package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"ldm/Initialize"
	"ldm/common/constant"
	"ldm/common/gen_proto/user"
	"ldm/services/user/srv"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(constant.RPC_USER_SRV),
	)
	// Initialise service
	service.Init()
	//初始化mysql
	db := Initialize.InitMysql()
	user.RegisterUserHandler(service.Server(), srv.NewUserImpl(db))
	// Run service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
