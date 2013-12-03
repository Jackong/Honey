/**
 * User: Jackong
 * Date: 13-11-24
 * Time: 下午5:13
 */
package net

import (
	"net"
	"io"
	"github.com/Jackong/log"
)


func SetUp(addr string, handler Handler, logger log.Logger) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Alert("accept|error|", err)
			continue
		}
		logger.Alert("accept|success|", conn.RemoteAddr())
		connection := NewConn(conn)
		Anonymous.Put(connection.Id, connection)
		go handleConn(connection, handler, logger)
	}
	return nil
}

func handleConn(conn *Conn, handler Handler, logger log.Logger) {
	defer func() {
		if e := recover(); e != nil {
			logger.Alert("handle|error|",e)
			if conn.IsSigned {
				Signed.Close(conn.Id)
			} else {
				Anonymous.Close(conn.Id)
			}
		}
	}()

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
	if _, err := io.ReadFull(conn, buf); err != nil {
		panic(err)
	}
}

func HandleWrite(conn net.Conn, buf []byte) {
	if _, err := conn.Write(buf); err != nil {
		panic(err)
	}
}
