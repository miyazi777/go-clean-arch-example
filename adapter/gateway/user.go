package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/miyazi777/go-clean-arch/entity"
	"github.com/miyazi777/go-clean-arch/usecase/port"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) port.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM `user` WHERE id=?", userID)
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found. UserID = %s", userID)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapter/gateway/GetUserByID")
	}

	return &user, nil
}

func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
