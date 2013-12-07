/**
 * User: Jackong
 * Date: 13-12-1
 * Time: 上午11:34
 */
package main

import (
	"github.com/Jackong/Honey/net"
	. "github.com/Jackong/Honey/global"
	"fmt"
	"github.com/Jackong/Honey/net/handler/json"
)

func main() {
	cln, err := net.NewClient("localhost" + Project.String("server", "addr"), &json.Handler{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write")
	request := json.NewRequest()
	request.Set("module", "signUp")
	request.Set("auth", "MTM4NjQxNjk0OXxmZkRsbmJfNHBITk0tQnZvUlBZU21NSGpvOTV0RUZJM3h3cUZMVWxrdkNVclJ6TVlpbWZKOWhwY2VWUm10UzQ2bXB3dmUyeWFyZTUyNndSd2VicGp8PjY_DlCA4QM5a5sc1ksyAWAs3oA5ppQ3iPdDPaRsX84=")
	request.Set("email", "jk7@gmail.com")
	request.Set("password", "01234567890123456789012345678901")
	request.Set("name", "jk6")
	err = cln.HandleWrite(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("read")
	response := json.NewResponse()
	cln.HandleRead(response)
	fmt.Println(response.Get("code"), response.Get("msg"), response.Get("auth"))

/*	request := json.NewRequest()
	request.Set("module", "checkAuth")
	request.Set("auth", `MTM4NjQxNjcyOXw2WjdkQWpaSFVuOW50NVA4cmY4azdER2U5RGlFMFZtU3NCMDJ3N052djBDaXdpcTJEdUtvMmNSUlRoenNBcUd4NHFLNkk1d2Jobkp3TWFfbmVYUGNkU2w0YUNHSVhUaVQ5cUFBcW03ckdTLXUzUUdiWWpRPXznb71q4KfUqxIzWJtsZ-ZMiuDCUOKMM8FyreYnrYb7xw==`)
	err = cln.HandleWrite(request)
	if err != nil {
		fmt.Println(err)
	}

	response := json.NewResponse()
	cln.HandleRead(response)
	fmt.Println(response.Get("code"), response.Get("msg"))*/
}

