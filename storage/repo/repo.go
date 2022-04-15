package repo

import (
	pu "tasks/Instagram_clone/insta_user/genproto/user_proto"
)

type UserStorageI interface {
	CreateUser(*pu.CreateUserReq) (*pu.GetUserRes, error)
	GetUser(*pu.GetUserReq) (*pu.GetUserRes, error)
	UpdateUser(*pu.UpdateUserReq) (*pu.GetUserRes, error)
	DeleteUser(*pu.DeleteUserReq) (*pu.Message, error)
	SearchUser(*pu.SearchUserReq) (*pu.UserList, error)
	Login(*pu.LoginReq) (*pu.GetUserRes, error)
}
