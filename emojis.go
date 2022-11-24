package main

import "time"

func tree() {
	disp.SetPixel(3, 1, white)

	disp.SetPixel(2, 2, white)
	disp.SetPixel(3, 2, white)
	disp.SetPixel(4, 2, white)

	disp.SetPixel(1, 3, white)
	disp.SetPixel(2, 3, white)
	disp.SetPixel(3, 3, white)
	disp.SetPixel(4, 3, white)
	disp.SetPixel(5, 3, white)

	disp.SetPixel(2, 4, white)
	disp.SetPixel(3, 4, white)
	disp.SetPixel(4, 4, white)

	disp.SetPixel(1, 5, white)
	disp.SetPixel(2, 5, white)
	disp.SetPixel(3, 5, white)
	disp.SetPixel(4, 5, white)
	disp.SetPixel(5, 5, white)

	disp.SetPixel(0, 6, white)
	disp.SetPixel(1, 6, white)
	disp.SetPixel(2, 6, white)
	disp.SetPixel(3, 6, white)
	disp.SetPixel(4, 6, white)
	disp.SetPixel(5, 6, white)
	disp.SetPixel(6, 6, white)

	disp.SetPixel(3, 7, white)

	disp.Display()
	time.Sleep(10 * time.Second)
	disp.ClearDisplay()
}

func snowflake(x, y int16) {
	disp.SetPixel(x, y-1, white)
	disp.SetPixel(x-1, y, white)
	disp.SetPixel(x, y, white)
	disp.SetPixel(x+1, y, white)
	disp.SetPixel(x, y+1, white)
}

func snowing() {
	for y := int16(0); y <= 14; y++ {
		snowflake(1, y-1)
		snowflake(6, y-4)
		snowflake(3, y-6)
		snowflake(5, y-10)
		snowflake(1, y-12)
		snowflake(6, y-14)
		disp.Display()
		time.Sleep(500 * time.Millisecond)
		disp.ClearDisplay()
	}
}

func smile() {
	disp.SetPixel(2, 0, white)
	disp.SetPixel(3, 0, white)
	disp.SetPixel(4, 0, white)
	disp.SetPixel(5, 0, white)

	disp.SetPixel(1, 1, white)
	disp.SetPixel(6, 1, white)

	disp.SetPixel(0, 2, white)
	disp.SetPixel(2, 2, white)
	disp.SetPixel(5, 2, white)
	disp.SetPixel(7, 2, white)

	disp.SetPixel(0, 3, white)
	disp.SetPixel(7, 3, white)

	disp.SetPixel(0, 4, white)
	disp.SetPixel(2, 4, white)
	disp.SetPixel(5, 4, white)
	disp.SetPixel(7, 4, white)

	disp.SetPixel(0, 5, white)
	disp.SetPixel(3, 5, white)
	disp.SetPixel(4, 5, white)
	disp.SetPixel(7, 5, white)

	disp.SetPixel(1, 6, white)
	disp.SetPixel(6, 6, white)

	disp.SetPixel(2, 7, white)
	disp.SetPixel(3, 7, white)
	disp.SetPixel(4, 7, white)
	disp.SetPixel(5, 7, white)

	disp.Display()
	time.Sleep(10 * time.Second)
	disp.ClearDisplay()
}

func heart() {
	disp.SetPixel(1, 0, white)
	disp.SetPixel(2, 0, white)
	disp.SetPixel(5, 0, white)
	disp.SetPixel(6, 0, white)

	disp.SetPixel(0, 1, white)
	disp.SetPixel(1, 1, white)
	disp.SetPixel(2, 1, white)
	disp.SetPixel(3, 1, white)
	disp.SetPixel(4, 1, white)
	disp.SetPixel(5, 1, white)
	disp.SetPixel(6, 1, white)
	disp.SetPixel(7, 1, white)

	disp.SetPixel(0, 2, white)
	disp.SetPixel(1, 2, white)
	disp.SetPixel(2, 2, white)
	disp.SetPixel(3, 2, white)
	disp.SetPixel(4, 2, white)
	disp.SetPixel(5, 2, white)
	disp.SetPixel(6, 2, white)
	disp.SetPixel(7, 2, white)

	disp.SetPixel(0, 3, white)
	disp.SetPixel(1, 3, white)
	disp.SetPixel(2, 3, white)
	disp.SetPixel(3, 3, white)
	disp.SetPixel(4, 3, white)
	disp.SetPixel(5, 3, white)
	disp.SetPixel(6, 3, white)
	disp.SetPixel(7, 3, white)

	disp.SetPixel(0, 4, white)
	disp.SetPixel(1, 4, white)
	disp.SetPixel(2, 4, white)
	disp.SetPixel(3, 4, white)
	disp.SetPixel(4, 4, white)
	disp.SetPixel(5, 4, white)
	disp.SetPixel(6, 4, white)
	disp.SetPixel(7, 4, white)

	disp.SetPixel(1, 5, white)
	disp.SetPixel(2, 5, white)
	disp.SetPixel(3, 5, white)
	disp.SetPixel(4, 5, white)
	disp.SetPixel(5, 5, white)
	disp.SetPixel(6, 5, white)

	disp.SetPixel(2, 6, white)
	disp.SetPixel(3, 6, white)
	disp.SetPixel(4, 6, white)
	disp.SetPixel(5, 6, white)

	disp.SetPixel(3, 7, white)
	disp.SetPixel(4, 7, white)

	disp.Display()
	time.Sleep(10 * time.Second)
	disp.ClearDisplay()
}

func dino() {
	disp.SetPixel(0, 0, white)
	disp.SetPixel(1, 0, white)
	disp.SetPixel(2, 0, white)
	disp.SetPixel(3, 0, white)

	disp.SetPixel(2, 1, white)
	disp.SetPixel(4, 1, white)

	disp.SetPixel(0, 2, white)
	disp.SetPixel(1, 2, white)
	disp.SetPixel(2, 2, white)
	disp.SetPixel(3, 2, white)
	disp.SetPixel(4, 2, white)
	disp.SetPixel(5, 2, white)

	disp.SetPixel(3, 3, white)
	disp.SetPixel(4, 3, white)
	disp.SetPixel(5, 3, white)

	disp.SetPixel(2, 4, white)
	disp.SetPixel(3, 4, white)
	disp.SetPixel(4, 4, white)
	disp.SetPixel(5, 4, white)

	disp.SetPixel(1, 5, white)
	disp.SetPixel(3, 5, white)
	disp.SetPixel(4, 5, white)
	disp.SetPixel(5, 5, white)
	disp.SetPixel(6, 5, white)

	disp.SetPixel(3, 6, white)
	disp.SetPixel(4, 6, white)
	disp.SetPixel(6, 6, white)

	disp.SetPixel(2, 7, white)
	disp.SetPixel(3, 7, white)
	disp.SetPixel(4, 7, white)
	disp.SetPixel(6, 7, white)
	disp.SetPixel(7, 7, white)

	disp.Display()
	time.Sleep(10 * time.Second)
	disp.ClearDisplay()
}