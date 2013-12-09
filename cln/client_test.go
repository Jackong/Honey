/**
 * User: Jackong
 * Date: 13-12-9
 * Time: 下午9:39
 */
package cln

import (
	"testing"
	"github.com/Jackong/Honey/net"
	"github.com/Jackong/Honey/net/handler/json"
	. "github.com/Jackong/Honey/global"
)

func client() *net.Client{
	cln, _ := net.NewClient("localhost" + Project.String("server", "addr"), &json.Handler{})
	return cln
}

func TestEmptyHeader(t *testing.T) {
	cln := client()
	if 	n, err := cln.Conn.Write([]byte("")); err != nil {
		t.Error(n, err)
	}
	buf := make([]byte, 12)
	_, err := cln.Conn.Read(buf)
	t.Log(err)
	if err == nil {
		t.Error("empty header, it should be close for time out")
	}
}


func TestSmallerHeader(t *testing.T) {
	cln := client()
	if 	n, err := cln.Conn.Write([]byte("1234567890")); err != nil {
		t.Error(n, err)
	}
	buf := make([]byte, 12)
	_, err := cln.Conn.Read(buf)
	t.Log(err)
	if err == nil {
		t.Error("not enough header, it should be close for time out")
	}
}

func TestBadHeader(t *testing.T) {
	cln := client()
	if n, err := cln.Conn.Write([]byte("123456789012")); err != nil {
		t.Error(n, err)
	}
	buf := make([]byte, 12)
	_, err := cln.Conn.Read(buf)
	t.Log(err)
	if err == nil {
		t.Error("client should be close for bad header")
	}
}

func TestLargeHeader(t *testing.T) {
	cln := client()
	if n, err := cln.Conn.Write([]byte(`{"len":10000}`)); err != nil {
		t.Error(n, err)
	}
	buf := make([]byte, 12)
	_, err := cln.Conn.Read(buf)
	t.Log(err)
	if err == nil {
		t.Error("so large header should be treated as bad header")
	}
}

func TestEmptyBody(t *testing.T) {
	cln := client()
	if n, err := cln.Conn.Write([]byte(`{"len":123}*`)); err != nil {
		t.Error(n, err)
	}
	buf := make([]byte, 12)
	_, err := cln.Conn.Read(buf)
	t.Log(err)
	if err == nil {
		t.Error("empty body, it should be close for time out")
	}
}

func TestBadSmallBody(t *testing.T) {
	cln := client()
	if n, err := cln.Conn.Write([]byte(`{"len":123}*{"module":"xxx"}`)); err != nil {
		t.Error(n, err)
	}
	buf := make([]byte, 12)
	_, err := cln.Conn.Read(buf)
	t.Log(err)
	if err == nil {
		t.Error("not enough body, it should be close for time out")
	}
}

func TestBadLargeBody(t *testing.T) {
	cln := client()
	if n, err := cln.Conn.Write([]byte(`{"len":12}*{"module":"xxxwwwwwww"}`)); err != nil {
		t.Error(n, err)
	}
	buf := make([]byte, 12)
	_, err := cln.Conn.Read(buf)
	t.Log(err)
	if err == nil {
		t.Error("so large body, it should be treated as bad body")
	}
}

func TestUnknownModule(t *testing.T) {
	request := json.NewRequest()
	request.Set("module", "xxx")
	moduleNotFound(request, t)
}

func TestUnsetModule(t *testing.T) {
	request := json.NewRequest()
	request.Set("notModule", "xxx")
	moduleNotFound(request, t)
}

func moduleNotFound(request net.Request, t *testing.T) {
	cln := client()
	response := json.NewResponse()
	cln.Handle(request, response)
	if request.Get("module") != response.Get("module") {
		t.Error("wrong response module")
	}
	if CODE_MODULE_NOT_FOUND != int(response.Get("code").(float64)) {
		t.Error("wrong response code", response.Get("code"))
	}
}
