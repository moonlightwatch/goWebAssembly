package main

import "goWebAssembly/functions"

func main() {
	done := make(chan bool, 1)
	functions.TestGet()
	go functions.MoveBall()
	<-done
}
