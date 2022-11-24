package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/tinyfont/proggy"

	"github.com/conejoninja/xmas-ball/max8x8display"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/drivers/max72xx"
)

var disp max8x8display.Device
var white = color.RGBA{255, 255, 255, 255}

func main() {
	// Pins for Arduino Nano 33 IOT
	err := machine.SPI0.Configure(machine.SPIConfig{
		SDO:       machine.D9, // default SDO pin
		SCK:       machine.D8, // default sck pin
		Frequency: 10000000,
	})

	if err != nil {
		println(err.Error())
	}

	matrix := max72xx.NewDevice(*machine.SPI0, machine.D7)

	disp = max8x8display.New(matrix)
	disp.Configure(max8x8display.Config{})

	disp.ClearDisplay()
	disp.Display()

	dino()
	heart()
	snowing()
	smile()
	tree()

	text := "TINYGO - MERRY XMAS"
	w32, _ := tinyfont.LineWidth(&proggy.TinySZ8pt7b, text)
	for {
		for i := int16(8); i > int16(-w32); i-- {
			disp.ClearDisplay()
			tinyfont.WriteLine(&disp, &proggy.TinySZ8pt7b, i, 7, text, white)
			disp.Display()
			time.Sleep(250 * time.Millisecond)
		}
	}
}
