package service

import (
	"context"
	pu "tasks/Instagram_clone/insta_user/genproto/user_proto"

	grpcClient "tasks/Instagram_clone/insta_user/service/grpc_client"
	"tasks/Instagram_clone/insta_user/storage"

	l "tasks/Instagram_clone/insta_user/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (r *PostService) CreateUser(ctx context.Context, req *pu.CreateUserReq) (*pu.GetUserRes, error) {
	res, err := r.storage.User().CreateUser(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (r *PostService) GetUser(ctx context.Context, req *pu.GetUserReq) (*pu.GetUserRes, error) {
	res, err := r.storage.User().GetUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) UpdateUser(ctx context.Context, req *pu.UpdateUserReq) (*pu.GetUserRes, error) {
	res, err := r.storage.User().UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) DeleteUser(ctx context.Context, req *pu.DeleteUserReq) (*pu.Message, error) {
	res, err := r.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) SearchUser(ctx context.Context, req *pu.SearchUserReq) (*pu.UserList, error) {
	res, err := r.storage.User().SearchUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) Login(ctx context.Context, req *pu.LoginReq) (*pu.GetUserRes, error) {
	res, err := r.storage.User().Login(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
