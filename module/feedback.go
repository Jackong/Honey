/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午10:48
 */
package module

import (
	"github.com/Jackong/Honey/net"
	"fmt"
)


func feedback(req net.Request, res net.Response, conn *net.Conn) error {
	fmt.Println(req.Get("module"))
	res.Set("code", 0)
	res.Set("msg", "hello daisy")
	return nil
}

func init() {
	net.AttachFunc("feedback", feedback)
}
