/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 下午4:32
 */
package funcs

import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
)

func beforeLog(req net.Request, res net.Response, conn *net.Conn) {
	Log.Infof("request|addr=%v,id=%v,signed=%v,module=%v", conn.Conn.RemoteAddr(), conn.Id, conn.IsSigned, req.Get("module"))
}

func afterSuccess(req net.Request, res net.Response, conn *net.Conn) {
	if net.HasResult(res) {
		return
	}
	net.Result(res, CODE_OK, "OK")
}


func init() {
	net.AllBefore(beforeLog)
	net.AllAfter(afterSuccess)
}
