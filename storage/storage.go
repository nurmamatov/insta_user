package storage

import (
	"tasks/Instagram_clone/insta_user/storage/postgres"
	"tasks/Instagram_clone/insta_user/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	User() repo.UserStorageI
}
type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}
