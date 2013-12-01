/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午10:20
 */
package net

type Module interface {
	Handle(Request, Response, *Conn)
}
