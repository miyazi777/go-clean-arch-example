package controller

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/miyazi777/go-clean-arch/usecase/port"
)

type User struct {
	// -> presenter.NewUserOutputPort
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort

	// -> interactor.NewUserInputPort
	InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort

	// -> gateway.NewUserRepository
	RepoFactory func(c *sql.DB) port.UserRepository

	Conn *sql.DB
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	outputPort := u.OutputFactory(w)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(ctx, userID)
}
