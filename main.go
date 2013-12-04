/**
 * User: Jackong
 * Date: 13-11-30
 * Time: 下午4:10
 */
package main

import (
	. "github.com/Jackong/Honey/global"
	"github.com/Jackong/Honey/net"
	"github.com/Jackong/Honey/net/handler/json"
	_ "github.com/Jackong/Honey/module/funcs"
	_ "github.com/Jackong/Honey/module"
	_ "github.com/Jackong/Honey/module/sign"
)

func main() {
	net.SetUp(Project.String("server", "addr"), &json.Handler{})
}
