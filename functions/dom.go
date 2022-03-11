package functions

import (
	"fmt"
	"syscall/js"
	"time"
)

func MoveBall() {
	ball := js.Global().Get("document").Call("getElementById", "ball")
	body := js.Global().Get("document").Get("body")
	ballStyle := ball.Get("style")
	ballStyle.Set("top", "0px")
	ballStyle.Set("left", "0px")
	location := js.Global().Get("document").Call("getElementById", "ball-location")

	top := 0
	left := 0
	directionX := 1
	directionY := 1
	t, _ := time.ParseDuration("0.005s")
	for {
		height := body.Get("clientHeight").Int()
		width := body.Get("clientWidth").Int()
		top += directionX
		left += directionY
		location.Set("innerText", fmt.Sprintf("[%d, %d]", left, top))
		ballStyle.Set("top", fmt.Sprintf("%dpx", top))
		ballStyle.Set("left", fmt.Sprintf("%dpx", left))

		if top > height-40 {
			directionX = -1
		}
		if left > width-40 {
			directionY = -1
		}
		if top < 0 {
			directionX = 1
		}
		if left < 0 {
			directionY = 1
		}
		time.Sleep(t)
	}
}
