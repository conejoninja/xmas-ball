package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/tinyfont/proggy"

	"github.com/conejoninja/xmas-ball/max8x8display"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/drivers/max72xx"
)

var disp max8x8display.Device
var white = color.RGBA{255, 255, 255, 255}
var msgs = []string{
	"HAPPY NEW YEAR",
	"HAPPY HOLIDAYS",
	"MERRY XMAS",
}
var w32 uint32
var n int32

func main() {
	// Let 1 second so all devices are powered correctly
	time.Sleep(1 * time.Second)

	// Initialize and configure the SPI connection
	err := machine.SPI0.Configure(machine.SPIConfig{
		SDO:       machine.D9, // default SDO pin
		SCK:       machine.D8, // default sck pin
		Frequency: 10000000,
	})

	if err != nil {
		println(err.Error())
	}

	// Create a new max7219 device, configure it as a 8x8 display
	matrix := max72xx.NewDevice(*machine.SPI0, machine.D7)
	disp = max8x8display.New(matrix)
	disp.Configure(max8x8display.Config{})

	// show nothing on the display, just in case
	disp.ClearDisplay()
	disp.Display()

	for {
		n = rand.Int31n(7)
		if n == 0 {
			heart() // shows a heart icon
		} else if n == 1 {
			snowing() // shows a small snowing animation
		} else if n == 2 {
			smile() // shows a smiley icon
		} else if n == 3 {
			tree() // shows a tree icon
		} else {
			// display one of the text messages in a marquee mode
			w32, _ = tinyfont.LineWidth(&proggy.TinySZ8pt7b, msgs[n-4])
			for i := int16(8); i > int16(-w32); i-- {
				disp.ClearDisplay()
				tinyfont.WriteLine(&disp, &proggy.TinySZ8pt7b, i, 7, msgs[n-4], white)
				disp.Display()
				time.Sleep(125 * time.Millisecond)
			}
		}
		disp.ClearDisplay()
		disp.Display()
		time.Sleep(5 * time.Second)
	}
}
