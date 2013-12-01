/**
 * User: Jackong
 * Date: 13-11-25
 * Time: 下午10:19
 */
package net
import (
	"net"
)

type Conn struct {
	Id string
	IsSigned bool
	net.Conn
}


func NewConn(conn net.Conn) *Conn{
	return &Conn{Id: conn.RemoteAddr().String(), Conn: conn}
}
