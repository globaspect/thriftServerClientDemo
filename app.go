package main

import "flag"

func main() {
	boolPtr := flag.Bool("server", true, "start as server")

	flag.Parse()

	if *boolPtr {
		Server()
	} else {
		Client()
	}
	
}
