package main

import (
	"fileStorage/localStorage"
)

func main() {
	err := localStorage.RunHttpFileServer("/home/lixiaohui/temp")
	if err != nil {
		panic(err)
	}
}
