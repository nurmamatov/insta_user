package db

import (
	"fmt"
	"tasks/Instagram_clone/insta_user/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)
	fmt.Println(psqlString)
	ConnDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return ConnDb, nil
}
