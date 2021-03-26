package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUrl(u string) (string, error) {
	u = "http://www.qq.com"
	res, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
	}
	r, err := ioutil.ReadAll(res)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(r)

	return u, nil
}
