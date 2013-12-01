/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午10:48
 */
package module

import (
	"fmt"
	"github.com/Jackong/Honey/net"
)

type feedback struct {

}

func (this *feedback) Handle(req net.Request, res net.Response, conn *net.Conn) {
	fmt.Println(req.Get("module"))
	res.Set("module", "feedback")
	res.Set("code", 0)
}

func init() {
	net.Attach("feedback", &feedback{})
}
