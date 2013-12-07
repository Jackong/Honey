/**
 * User: Jackong
 * Date: 13-12-2
 * Time: 下午10:52
 */
package sign

import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
	"github.com/Jackong/Honey/service"
)

const (
	RE_EMAIL = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	RE_PASSWORD = `[0-9a-f]{32}`
)

func signUp(request net.Request, response net.Response, conn *net.Conn) (err error) {
	email := net.Pattern(request, "email", RE_EMAIL)
	password := net.Pattern(request, "password", RE_PASSWORD)
	name := net.Pattern(request, "name", `[0-9a-zA-Z]{2,15}`)

	if !service.User.SignUp(email, password, name) {
		net.Result(response, CODE_FAILED, "Account is exists")
	}
	return
}

func signIn(request net.Request, response net.Response, conn *net.Conn) (err error) {
	email := net.Pattern(request, "email", RE_EMAIL)
	password := net.Pattern(request, "password", RE_PASSWORD)

	if !service.User.SignIn(email, password) {
		net.Result(response, CODE_FAILED, "Account or password is wrong")
	}
	return
}

func init() {
	net.AttachFunc("signUp", signUp)
	net.AttachFunc("signIn", signIn)
}
