package entity

import (
	"mymicrosystem/src/share/pb"
)

type User struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (u User) ToProtoUser() *pb.User {
	return &pb.User{
		Id:      u.Id,
		Name:    u.Name,
		Address: u.Address,
		Phone:   u.Phone,
	}
}
