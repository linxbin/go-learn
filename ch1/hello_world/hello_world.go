package main

import "os"

func main() {
	if len(os.Args) > 1 {
		print("hello world ", os.Args[1])
	} else {
		print("hello world none args")
	}

}
