package main

import (
	"fileStorage/localStorage"
)

func main() {
	err := localStorage.RunHttpFileServer("/home/cyx1/Desktop/tmp")
	if err != nil {
		panic(err)
	}
}
