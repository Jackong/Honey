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
	"fmt"
)

func signUp(req net.Request, res net.Response, conn *net.Conn) error {
	user := Col("user")
	email := req.Get("email")
	password := req.Get("password")
	name := req.Get("name")
	fmt.Println(email, password, name)
	_, err := user.Find(db.Cond{
		"id": email,
	})
	if err == nil {
		res.Set("code", 2)
		res.Set("tips", "the account is exists")
		return nil
	}
	userMeta := &meta.User{Password: proto.String(password.(string)), Name: proto.String(name.(string))}
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
	res.Set("code", 0)
	return nil
}

func init() {
	net.AttachFunc("signUp", signUp)
}
