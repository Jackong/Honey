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
)

var (
	User user
)

func init() {
	User = user{user: model.NewModel("user")}
}

type user struct {
	user *model.Model
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
