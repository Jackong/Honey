/**
 * User: Jackong
 * Date: 13-12-7
 * Time: 上午9:35
 */
package service

import (
	. "github.com/Jackong/Honey/global"
	"github.com/Jackong/Honey/net"
	"github.com/gorilla/securecookie"
	"menteslibres.net/gosexy/db"
	"github.com/Jackong/Honey/model"
)

var (
	User user
)

func init() {
	User = user{
		user: model.Collection("user"),
		auth: securecookie.New([]byte(Project.String("auth", "hash")), []byte(Project.String("auth", "block"))),
	}
}

type user struct {
	user db.Collection
	auth *securecookie.SecureCookie
}

func (this user) SignUp(email, password, name string) bool {
	item, err := this.user.Find(db.Cond{
		"email": email,
	})
	if err == nil && item != nil {
		Log.Errorf("Account is exists:%v", email)
		return false
	}
	_, err = this.user.Append(db.Item{
		"email": email,
		"password": password,
		"name": name,
	})
	if err != nil {
		Log.Errorf("%v|Append meta:%v", this.user.Name(), err)
		return false
	}
	return true
}

func (this user) SignIn(email, password string) bool {
	item, err := this.user.Find(db.Cond{
		"email": email,
	})
	if err != nil {
		Log.Errorf("%v|%v is not exists:%v", this.user.Name(), email, err)
		return false
	}

	return item["password"] == password
}

func (this user) Auth(conn *net.Conn, email string) (string, bool) {
	conn.Info["id"] = email
	conn.Info["authTime"] = TimeStamp()
	auth, err := this.auth.Encode("auth", conn.Info)
	if err != nil {
		Log.Alertf("auth|%v|Can not encode auth:%v", conn.Id, err)
		return auth, false
	}

	net.SignIn(conn, email)
	return auth, true
}

func (this user) IsAuth(request net.Protocol, conn *net.Conn) bool {
	auth := net.Required(request, "auth")
	if err := this.auth.Decode("auth", auth.(string), &conn.Info); err != nil {
		Log.Alertf("isAuth|%v|Can not decode auth:%v|%v", conn.Id, auth, err)
		return false
	}

	id, ok := conn.Info["id"]
	if !ok {
		Log.Alertf("isAuth|%v|Auth must be include id:%v", conn.Id, conn.Info)
		return false
	}
	return net.SignIn(conn, id.(string))
}
