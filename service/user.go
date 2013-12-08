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
	"labix.org/v2/mgo/bson"
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

func (this user) SignUp(conn *net.Conn, email, password, name string) (auth string, ok bool) {
	item, err := this.user.Find(db.Cond{
		"email": email,
	})
	if err != nil {
		Log.Errorf("Find by email failed:%v|%v", email, err)
		return
	}

	if item != nil {
		Log.Errorf("Account is exists:%v", email)
		return
	}
	item = db.Item{
		"email": email,
		"password": password,
		"name": name,
	}
	ids, err := this.user.Append(item)
	if err != nil {
		Log.Errorf("%v|Append meta:%v", this.user.Name(), err)
		return
	}
	return this.Auth(conn, string(ids[0]), item)
}

func (this user) SignIn(conn *net.Conn, email, password string) (auth string, ok bool) {
	item, err := this.user.Find(db.Cond{
		"email": email,
	})
	if err != nil {
		Log.Errorf("%v|%v is not exists:%v", this.user.Name(), email, err)
		return
	}

	if item["password"] != password {
		return
	}

	return this.Auth(conn, item["_id"].(bson.ObjectId).Hex(), item)
}

func (this user) Auth(conn *net.Conn, id string, item db.Item) (string, bool) {
	conn.Info["id"] = id
	conn.Info["email"] = item["email"]
	conn.Info["authTime"] = TimeStamp()
	auth, err := this.auth.Encode("auth", conn.Info)
	if err != nil {
		Log.Alertf("auth|%v|Can not encode auth:%v", conn.Id, err)
		return auth, false
	}

	return auth, net.SignIn(conn, id)
}

func (this user) CheckAuth(request net.Protocol, conn *net.Conn) bool {
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
	if _, ok = conn.Info["email"]; !ok {
		Log.Alertf("checkAuth|%v|Auth must be include email:%v", conn.Id, conn.Info)
		return false
	}
	return net.SignIn(conn, id.(string))
}
