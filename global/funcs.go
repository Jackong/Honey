/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 下午4:32
 */
package global

import (
	"github.com/Jackong/Honey/net"
)

func beforeLog(req net.Request, res net.Response, conn *net.Conn) error {
	Log.Infof("request|addr=%v,id=%v,signed=%v,module=%v", conn.Conn.RemoteAddr(), conn.Id, conn.IsSigned, req.Get("module"))
	return nil
}

func afterSuccess(req net.Request, res net.Response, conn *net.Conn) error {
	if res.Get("code") != nil {
		return nil
	}
	res.Set("code", 0)
	return nil
}


func init() {
	net.AllBefore(beforeLog)
	net.AllAfter(afterSuccess)
}
