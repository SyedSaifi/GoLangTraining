package main

import (
	"os"
)
func main() {

	//panic("big problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}