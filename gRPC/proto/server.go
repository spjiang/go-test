package proto

import (
	"context"
	"fmt"
)

// 定义空接口
type UserInfoService struct{}

// 定义一个接口类型
var U = UserInfoService{}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *UserRequest) (resp *UserResponse, err error) {
	// 通过用户名查询用户信息
	name := req.Name
	fmt.Println(name)
	// 数据库查询用户信息
	if name == "ZS" {
		resp = &UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"SING", "RUN"},
		}
	}
	return
}
func (s *UserInfoService) mustEmbedUnimplementedUserInfoServiceServer() {}
