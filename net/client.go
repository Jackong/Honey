/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午11:21
 */
package net

import (
	"net"
)

type client struct {
	net.Conn
	handler Handler
}

func NewClient(addr string, handler Handler) (cln *client, err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return cln, err
	}
	cln = &client{Conn: conn, handler: handler}
	return cln, nil
}


func (this *client) HandleWrite(req Request) error {
	buf, err := this.handler.FormatProtocol(req)
	if err != nil {
		return err
	}
	this.Conn.Write(buf)
	return nil
}

func (this *client) HandleRead(res Response) {
	header := make([]byte, this.handler.HeaderLength())
	HandleRead(this.Conn, header)
	length, err := this.handler.HandleHeader(header)
	if err != nil {
		panic(err)
	}
	response := make([]byte, length)
	HandleRead(this.Conn, response)
	err = res.Decode(response)
	if err != nil {
		panic(err)
	}
}
