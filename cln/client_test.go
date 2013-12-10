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
	"strings"
)

func client() *net.Client{
	cln, _ := net.NewClient("localhost" + Project.String("server", "addr"), &json.Handler{})
	return cln
}

var (
	customProtocolTest = []struct {
	input string
	error string
}{
	{"", "EOF"},
	{"1234567890", "EOF"},
	{"123456789012", "EOF"},
	{`{"len":10000}`, "The specified network name is no longer available."},
	{`{"len":123}*`, "EOF"},
	{`{"len":123}*{"module":"xxx"}`, "EOF"},
	{`{"len":13}**{"module":""}`, "EOF"},
	{`{"len":14}**{"module":"x"}`, "EOF"},
	{`{"len":15}**{"module":null}`, "EOF"},
	{`{"len":15}**{"module":1.25}`, "EOF"},
	{`{"len":12}**{"module":0}`, "EOF"},
	{`{"len":15}**{"module":true}`, "EOF"},
	{`{"len":19}**{"notModule":"xxx"}`, "EOF"},
	{`{"len":12}*{"module":"xxxwwwwwww"}`, "The specified network name is no longer available."},
}
	moduleNotFoundTest = []struct {
	key string
	value interface {}
	code int
}{
	{"module", "xx", CODE_MODULE_NOT_FOUND},
}
)

func TestCustomProtocol(t *testing.T) {
	for _, data := range customProtocolTest {
		cln := client()
		if 	n, err := cln.Conn.Write([]byte(data.input)); err != nil {
			t.Error(n, err)
		}
		buf := make([]byte, 12)
		_, err := cln.Conn.Read(buf)

		if !strings.Contains(err.Error(), data.error) {
			t.Errorf("cln.Write(%v): expect [%v], actual [%v]", data.input, data.error, err.Error())
		}
	}
}

func TestModuleNotFound(t *testing.T) {
	for _, data := range moduleNotFoundTest {
		cln := client()
		request, response := json.NewRequest(), json.NewResponse()
		request.Set(data.key, data.value)
		cln.Handle(request, response)
		if request.Get("module") != response.Get("module") {
			t.Errorf("%v:%v|response.Get(\"module\"): expect [%v], actual [%v]", data.key, data.value, request.Get("module"), response.Get("module"))
		}
		if data.code != int(response.Get("code").(float64)) {
			t.Errorf("%v:%v|response.Get(\"code\"): expect [%v], actual [%v]", data.key, data.value, data.code, response.Get("code"))
		}
	}
}
