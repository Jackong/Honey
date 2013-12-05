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
	e "github.com/Jackong/Honey/err"
)

const (
	RE_EMAIL = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	RE_PASSWORD = `[0-9a-f]{32}`
)

func signUp(req net.Request, res net.Response, conn *net.Conn) error {
	email := net.Pattern(req, "email", RE_EMAIL)
	password := net.Pattern(req, "password", RE_PASSWORD)
	name := net.Required(req, "name")

	user := Col("user")
	_, err := user.Find(db.Cond{
		"id": email,
	})
	if err == nil {
		panic(e.Runtime{Code: e.CODE_USER_ACCOUNT_EXIST, Msg: "the account is exists"})
	}
	userMeta := &meta.User{Password: proto.String(password), Name: proto.String(name.(string))}
	buf, err := proto.Marshal(userMeta)
	if err != nil {
		return err
	}
	_, err = user.Append(db.Item{
		"id": email,
		"data": buf,
	})
	if err != nil {
		return err
	}
	return nil
}

func init() {
	net.AttachFunc("signUp", signUp)
}
