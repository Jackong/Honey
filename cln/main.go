/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午11:34
 */
package cln

import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
	"fmt"
	"github.com/Jackong/Honey/net/handler/json"
)

func main() {
	cln, _ := net.NewClient("localhost" + Project.String("server", "addr"), &json.Handler{})
	request := json.NewRequest()
	request.Set("module", "signIn")
	request.Set("email", "jk2@gmail.com")
	request.Set("password", "01234567890123456789012345678901")
	request.Set("name", "jk6")
	response := json.NewResponse()
	cln.Handle(request, response)
	fmt.Println(response.Get("code"), response.Get("msg"), response.Get("auth"))
}

