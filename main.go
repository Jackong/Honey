/**
 * User: Jackong
 * Date: 13-11-30
 * Time: 下午4:10
 */
package main

import (
	. "github.com/Jackong/Honey/global"
	"github.com/Jackong/Honey/net"
	_ "github.com/Jackong/Honey/module"
)

func main() {
	net.SetUp(Project.String("server", "addr"), Handler, Log)
}
