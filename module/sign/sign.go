/**
 * User: Jackong
 * Date: 13-12-2
 * Time: 下午10:52
 */
package sign

import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
	"github.com/Jackong/db"
	"code.google.com/p/goprotobuf/proto"
	"github.com/Jackong/Honey/meta"
)

const (
	RE_EMAIL = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	RE_PASSWORD = `[0-9a-f]{32}`
)

func signUp(request net.Request, response net.Response, conn *net.Conn) (err error) {
	email := net.Pattern(request, "email", RE_EMAIL)
	password := net.Pattern(request, "password", RE_PASSWORD)
	name := net.Required(request, "name")

	user := Col("user")
	_, err = user.Find(db.Cond{
		"id": email,
	})
	if err == nil {
		net.Result(response, CODE_USER_ACCOUNT_EXIST, "This account is exists")
		return
	}
	userMeta := &meta.User{Password: proto.String(password), Name: proto.String(name.(string))}
	buf, err := proto.Marshal(userMeta)
	if err != nil {
		net.Result(response, CODE_FAILED, "Sign up failed", err)
		return
	}
	_, err = user.Append(db.Item{
		"id": email,
		"data": buf,
	})
	if err != nil {
		net.Result(response, CODE_FAILED, "Sign up failed", err)
	}
	return
}

func init() {
	net.AttachFunc("signUp", signUp)
}
