package postgres

import (
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	pu "tasks/Instagram_clone/insta_user/genproto/user_proto"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}
func (r *UserRepo) CreateUser(req *pu.CreateUserReq) (*pu.GetUserRes, error) {
	var (
		UserId string
		Now    string
	)
	queryUser := `INSERT INTO users (
				user_id, 
				firs_name,
				last_name,
				username,
				password,
				phone,
				email,
				gender,
				created_at)
			VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	queryPhoto := `INSERT INTO user_photo (image_id, user_id, type, base_code) VALUES($1,$2,$3,$4)`

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("Erro while begin tx", err)
		return nil, err
	}

	UserId = uuid.New().String()
	Now = time.Now().Format(time.RFC3339)

	_, err = tx.Exec(queryUser, UserId, req.FirstName, req.LastName, req.Username, req.Password, req.Phone, req.Email, req.Gender, Now)
	if err != nil {
		log.Println("Erro while insert user", err)
		tx.Rollback()
		return nil, err
	}

	img := strings.Split(req.Photo, ",")
	_, err = tx.Exec(queryPhoto, UserId, img[0], img[1])
	if err != nil {
		log.Println("Erro while insert user_photo", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return r.GetUser(&pu.GetUserReq{Username: req.Username})
}
func (r *UserRepo) GetUser(req *pu.GetUserReq) (*pu.GetUserRes, error) {
	var (
		ImgType  string
		BaseCode string
	)
	res := pu.GetUserRes{}
	queryUser := `SELECT user_id, first_name, last_name, username, phone, email, gender, created_at FROM users WHERE username=$1 AND deleted_at IS NULL`
	queryPhoto := `SELECT type, base_code FROM user_photo WHERE user_id=$1`
	err := r.db.QueryRow(queryUser, req.Username).Scan(
		&res.UserId,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.Phone,
		&res.Email,
		&res.Gender,
		&res.CreatedAt,
	)
	if err != nil {
		log.Println("Error while get user", err)
		return nil, err
	}

	err = r.db.QueryRow(queryPhoto, res.UserId).Scan(
		&ImgType,
		&BaseCode,
	)
	if err != nil {
		log.Println("Error while get user photo", err)
		return nil, err
	}
	res.Photo = ImgType + BaseCode
	return &res, nil
}
func (r *UserRepo) UpdateUser(req *pu.UpdateUserReq) (*pu.GetUserRes, error) {
	queryUser := `UPDATE users SET first_name=$2, last_name=$3, username=$4, phone=$5, email=$6, gender=$7 WHERE user_id=$1 AND deleted_at IS NULL`
	queryPhoto := `UPDATE user_photo SET type=$2, base_code=$3 WHERE user_id=$1`

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("Error while begin tx in update user", err)
		return nil, err
	}
	_, err = tx.Exec(queryUser, req.UserId, req.FirstName, req.LastName, req.Username, req.Phone, req.Email, req.Gender)
	if err != nil {
		log.Println("Error while update user", err)
		return nil, err
	}
	img := strings.Split(req.Photo, ",")
	_, err = tx.Exec(queryPhoto, req.UserId, img[0], img[1])
	if err != nil {
		log.Println("Error while update user_photo", err)
		return nil, err
	}
	return r.GetUser(&pu.GetUserReq{Username: req.Username})
}
func (r *UserRepo) DeleteUser(req *pu.DeleteUserReq) (*pu.Message, error) {
	queryUser := `UPDATE users SET deleted_at=$2 WHERE user_id=$1 AND deleted_a IS NULL`
	queryPhoto := `DELETE FROM user_photo WHERE user_id=$1`

	now := time.Now().Format(time.RFC3339)
	_, err := r.db.Exec(queryUser, req.UserId, now)
	if err != nil {
		log.Println("Error while delete user")
		return nil, err
	}

	_, err = r.db.Exec(queryPhoto, req.UserId)
	if err != nil {
		log.Println("Error while delete user_photo")
		return nil, err
	}

	return &pu.Message{Message: "Deleted!"}, nil
}
func (r *UserRepo) SearchUser(req *pu.SearchUserReq) (*pu.UserList, error) {
	res := pu.UserList{}
	query := `SELECT user_id, username FROM users WHERE username LIKE '$1%'`
	rows, err := r.db.Query(query, req.Username)
	if err != nil {
		log.Println("Erro while search")
		return nil, err
	}
	for rows.Next() {
		result := pu.SearchRes{}
		err = rows.Scan(
			&result.UserId,
			&result.Username,
		)
		if err != nil {
			log.Println("Error while get users in search for")
			return nil, err
		}
		res.Users = append(res.Users, &result)
	}

	return &res, nil
}

func (r *UserRepo) Login(req *pu.LoginReq) (*pu.GetUserRes, error) {

	var (
		check int
	)
	query := `SELECT COUNT(*) FROM users WHERE username=$1 AND password=$2 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, req.Username, req.Password).Scan(
		&check,
	)
	if err != nil {
		log.Println("Error while login")
		return nil, err
	}

	return r.GetUser(&pu.GetUserReq{Username: req.Username})
}
func (r *UserRepo) UpdatePassword(req *pu.UpdatePass) (*pu.Message, error) {
	return nil, nil
}
