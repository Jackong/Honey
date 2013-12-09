/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午11:21
 */
package net

import (
	"net"
	. "github.com/Jackong/Honey/global"
)

type Client struct {
	net.Conn
	handler Handler
}

func NewClient(addr string, handler Handler) (cln *Client, err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		Log.Fatalf("new Client:%v", err)
		return
	}
	cln = &Client{Conn: conn, handler: handler}
	return
}

func (this *Client) Handle(request Request, response Response) {
	defer func() {
		if e := recover(); e != nil {
			Log.Fatalf("Client|handle:%v", e)
		}
	}()
	this.HandleWrite(request)
	this.HandleRead(response)
}

func (this *Client) HandleWrite(request Request) {
	buf, err := this.handler.FormatProtocol(request)
	if err != nil {
		panic(err)
	}

	for n := 0; n < len(buf); {
		i, err := this.Conn.Write(buf[n:])
		if err != nil {
			panic(err)
		}
		n+=i
	}
}

func (this *Client) HandleRead(response Response) {
	header := make([]byte, this.handler.HeaderLength())
	HandleRead(this.Conn, header)
	length, err := this.handler.HandleHeader(header)
	if err != nil {
		panic(err)
	}
	resBuf := make([]byte, length)
	HandleRead(this.Conn, resBuf)
	err = response.Decode(resBuf)
	if err != nil {
		panic(err)
	}
}
