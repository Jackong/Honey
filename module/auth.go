/**
 * User: Jackong
 * Date: 13-12-7
 * Time: 下午5:39
 */
package module


import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
	"github.com/Jackong/Honey/service"
)

func checkAuth(request net.Request, response net.Response, conn *net.Conn) (err error) {
	if !service.User.IsAuth(request, conn) {
		net.Result(response, CODE_FAILED, "Auth failed")
	}
	return
}

func init() {
	net.AttachFunc("checkAuth", checkAuth)
}


