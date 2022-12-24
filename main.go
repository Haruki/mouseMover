package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	//centerX := 500.0
	//centerY := 500.0
	//radius := 200

	for {
		robotgo.Move(100, 100)
		time.Sleep(time.Second * 10)
		robotgo.Move(200, 200)
		/*
			angle := time.Since(time.Now()).Seconds() / 60 * 2 * math.Pi
			x := int(centerX + float64(radius)*math.Cos(angle))
			y := int(centerY + float64(radius)*math.Sin(angle))
			fmt.Printf("Moving mouse to (%d, %d)\n", x, y)
			robotgo.Move(x, y)
		*/
		time.Sleep(time.Second)
	}
}
