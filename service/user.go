/**
 * User: Jackong
 * Date: 13-12-7
 * Time: 上午9:35
 */
package service

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/Jackong/Honey/meta"
	. "github.com/Jackong/Honey/global"
	"github.com/Jackong/Honey/model"
	"github.com/Jackong/Honey/net"
	"github.com/gorilla/securecookie"
)

var (
	User user
)

func init() {
	User = user{
		user: model.NewModel("user"),
		auth: securecookie.New([]byte(Project.String("auth", "hash")), []byte(Project.String("auth", "block"))),
	}
}

type user struct {
	user *model.Model
	auth *securecookie.SecureCookie
}

func (this user) SignUp(email, password, name string) bool {
	if this.user.Exist(email) {
		Log.Errorf("Account is exists:%v", email)
		return false
	}
	userMeta := &meta.User{Password: proto.String(password), Name: proto.String(name)}
	return this.user.Set(email, userMeta)
}

func (this user) SignIn(email, password string) bool {
	userMeta := &meta.User{}
	if !this.user.Get(email, userMeta) {
		return false
	}
	return userMeta.GetPassword() == password
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
