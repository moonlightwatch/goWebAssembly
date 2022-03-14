package functions

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

var Balls []string
var r1 *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func AddBall() string {
	// <span id="ball" style="position: absolute;border-radius: 20px; width:40px; height: 40px; background-color: blueviolet;"></span>
	// document.body.innerHTML += "<h1>666</h1>";
	id := fmt.Sprintf("ball_%d", len(Balls))
	body := js.Global().Get("document").Get("body")
	body.Set(
		"innerHTML",
		fmt.Sprintf(
			"%s\n%s",
			body.Get("innerHTML").String(),
			fmt.Sprintf(`<span id="%s" style="position: absolute;border-radius: 5px; width:10px; height: 10px; background-color: blueviolet;"></span>`, id)),
	)
	Balls = append(Balls, id)
	return id
}

func MoveBall(id string) {
	ball := js.Global().Get("document").Call("getElementById", id)
	body := js.Global().Get("document").Get("body")
	ballStyle := ball.Get("style")
	ballStyle.Set("top", "0px")
	ballStyle.Set("left", "0px")

	top := 0
	left := 0
	directionX := 1
	directionY := 1
	t, _ := time.ParseDuration("0.002s")
	for {
		height := body.Get("clientHeight").Int()
		width := body.Get("clientWidth").Int()
		top, _ = strconv.Atoi(strings.Replace(ballStyle.Get("top").String(), "px", "", -1))
		top += directionX
		left, _ = strconv.Atoi(strings.Replace(ballStyle.Get("left").String(), "px", "", -1))
		left += directionY
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

func RandomBall(id string) {
	ball := js.Global().Get("document").Call("getElementById", id)
	body := js.Global().Get("document").Get("body")
	ballStyle := ball.Get("style")
	ballStyle.Set("top", "0px")
	ballStyle.Set("left", "0px")
	top := 0
	left := 0
	t, _ := time.ParseDuration("2s")

	for {
		time.Sleep(t)
		top = r1.Intn(body.Get("clientHeight").Int() - 40)
		left = r1.Intn(body.Get("clientWidth").Int() - 40)
		ballStyle.Set("top", fmt.Sprintf("%dpx", top))
		ballStyle.Set("left", fmt.Sprintf("%dpx", left))
	}
}
