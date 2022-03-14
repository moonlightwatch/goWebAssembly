package main

import (
	"goWebAssembly/functions"
	"time"
)

func main() {
	done := make(chan bool, 1)
	functions.TestGet()
	t, _ := time.ParseDuration("0.5s")
	id := functions.AddBall()
	go functions.MoveBall(id)
	go functions.RandomBall(id)
	time.Sleep(t)
	go functions.RandomBall(id)
	<-done
}
