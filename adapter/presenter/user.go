package presenter

import (
	"fmt"
	"net/http"

	"github.com/miyazi777/go-clean-arch/entity"
	"github.com/miyazi777/go-clean-arch/usecase/port"
)

type User struct {
	w http.ResponseWriter
}

func NewUserOutputort(w http.ResponseWriter) port.UserOutputPort {
	return &User{w: w}
}

// 正常時のHTTP Responseをここでrenderしている
func (u *User) Render(user *entity.User) {
	u.w.WriteHeader(http.StatusOK)
	fmt.Fprint(u.w, user.Name)
}

// 異常時のHTTP Responseをここでrenderしている
func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}
