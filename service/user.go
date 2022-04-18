package service

import (
	"context"
	"fmt"
	pp "tasks/Instagram_clone/insta_user/genproto/post_proto"
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
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	list, err := r.client.PostService().ListUserPosts(ctx, &pp.ListPostsReq{UserId: res.UserId})
	if err != nil {
		return nil, err
	}
	for _, i := range list.Posts {
		posts := pu.Posts{
			UserId:      i.UserId,
			PostId:      i.PostId,
			CheckLike:   i.CheckLike,
			Title:       i.Title,
			Description: i.Description,
			Image:       i.Image,
			Likes:       i.Likes,
			CreatedAt:   i.CreatedAt,
		}
		for _, j := range i.Comments {
			comments := pu.Comment{
				CommentId: j.CommentId,
				UserId:    j.UserId,
				Text:      j.Text,
			}
			posts.Comments = append(posts.Comments, &comments)
		}
		res.Posts = append(res.Posts, &posts)
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
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return r.GetUser(ctx, &pu.GetUserReq{Username: res})
}
func (r *PostService) UpdatePassword(ctx context.Context, req *pu.UpdatePass) (*pu.Message, error) {
	res, err := r.storage.User().UpdatePassword(req.UserId, req.NewPassword, req.OldPassword)
	if err != nil {
		return nil, err
	}
	return res, nil
}
