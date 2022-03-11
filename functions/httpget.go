package functions

import (
	"fmt"
	"io"
	"net/http"
)

// Get 可以实现异步请求（如同 ajax），但是要注意跨域问题
func Get(path string) (string, error) {
	r, err := http.DefaultClient.Get(path)
	if err != nil {
		return "", err
	} else {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return "", err
		} else {
			return string(b), nil
		}
	}
}

func TestGet() {
	respBody, err := Get("/run.sh")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(respBody)
	}
}
