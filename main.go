package main

import (
	"fmt"
	"ziweiSgServer/config"
)

func main() {
	host := config.File.MustValue("login_server", "host", "111")
	fmt.Println(host)
}
