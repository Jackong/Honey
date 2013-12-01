/**
 * User: Jackong
 * Date: 13-11-30
 * Time: 下午9:44
 */
package net

type Handler interface {
	HandleAcceptError(error)
	HandleConnError(interface {})
	//get header length
	HeaderLength() int
	//pass header buffer and return the length of request
	HandleHeader([]byte) (uint, error)
	//pass the request buffer and return the response buffer
	HandleRequest([]byte, *Conn) ([]byte, error)
	//format protocol by header
	FormatProtocol(protocol Protocol)([]byte, error)
}
