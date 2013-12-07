/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午10:48
 */
package module

import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
)


func feedback(request net.Request, response net.Response, conn *net.Conn) error {
	contact := request.Get("contact")
	content := net.Required(request, "content")
	Log.Infof("feedback|%v,%v,%v,%v:%v", conn.Id, conn.IsSigned, conn.Info, contact, content)
	return nil
}

func init() {
	net.AttachFunc("feedback", feedback)
}
