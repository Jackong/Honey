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


func feedback(req net.Request, res net.Response, conn *net.Conn) error {
	contact := net.Required(req, "contact")
	content := net.Required(req, "content")
	Log.Infof("feedback|%v,%v:%v", conn.Id, contact, content)
	return nil
}

func init() {
	net.AttachFunc("feedback", feedback)
}
