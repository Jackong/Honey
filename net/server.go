/**
 * User: Jackong
 * Date: 13-11-24
 * Time: 下午5:13
 */
package net

import (
	"net"
	"io"
	. "github.com/Jackong/Honey/global"
	"time"
)

var (
	readDeadline time.Duration
	writeDeadline time.Duration
)

func init() {
	readDeadline = time.Duration(Project.Get("deadline", "read").(float64) * 1e9)
	writeDeadline = time.Duration(Project.Get("deadline", "write").(float64) * 1e9)
}

func SetUp(addr string, handler Handler) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			Log.Alert("accept|error|", err)
			continue
		}
		Log.Alert("accept|success|", conn.RemoteAddr())
		connection := NewConn(conn)
		Anonymous.Put(connection.Id, connection)
		go handleConn(connection, handler)
	}
	return nil
}

func handleConn(conn *Conn, handler Handler) {
	defer func() {
		if e := recover(); e != nil {
			Log.Alert("handle|error|",e)
			Close(conn)
		}
	}()

	//panic when read or write error
	for {
		header := make([]byte, handler.HeaderLength())
		HandleRead(conn, header)
		length, err := handler.HandleHeader(header)
		if err != nil {
			panic(err)
		}
		request := make([]byte, length)
		HandleRead(conn, request)
		response, err := handler.HandleRequest(request, conn)
		if err != nil {
			panic(err)
		}
		HandleWrite(conn, response)
	}
}

func HandleRead(conn net.Conn, buf []byte) {
	conn.SetReadDeadline(Time().Add(readDeadline))
	if _, err := io.ReadFull(conn, buf); err != nil {
		panic(err)
	}
}

func HandleWrite(conn net.Conn, buf []byte) {
	conn.SetReadDeadline(Time().Add(writeDeadline))
	for n := 0; n < len(buf); {
		if i, err := conn.Write(buf[n:]); err != nil {
			panic(err)
			n += i
		}
	}
}
