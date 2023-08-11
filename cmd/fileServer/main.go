package main

import (
	"fileStorage/localStorage"
)

func main() {
	err := localStorage.RunHttpFileServer("/home/navy/Desktop/tmp")
	if err != nil {
		panic(err)
	}
}
