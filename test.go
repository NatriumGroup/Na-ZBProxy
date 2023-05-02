package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
   /* 这是我的第一个简单的程序 */
   res, err := http.Get("http://124.223.44.130/api/v2/ZBP.php?ign=China__mxb")
if err != nil {
	return
}
defer res.Body.Close()
body, _ := ioutil.ReadAll(res.Body)
//bodystr := string(body)


if string(body)==string(`"200"`) {
	fmt.Println("Hello, World!")
}else{
	fmt.Println(string(body))
	fmt.Println("200")
}
}