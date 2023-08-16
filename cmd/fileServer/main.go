package main

import (
	"fileStorage/localStorage"
)

func main() {
	err := localStorage.RunHttpFileServer("/home/navy/Desktop/temp")
	if err != nil {
		panic(err)
	}
}
