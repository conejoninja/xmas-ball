package max8x8display

import (
	"image/color"
	"tinygo.org/x/drivers/max72xx"
)

const (
	NO_ROTATION  Rotation = 0
	ROTATION_90  Rotation = 1 // 90 degrees clock-wise rotation
	ROTATION_180 Rotation = 2
	ROTATION_270 Rotation = 3
)

type Rotation uint8

// Device holds LEDStriper device and some other information
type Device struct {
	matrix   *max72xx.Device
	rotation Rotation
	buffer   [8]byte
}

// Config is the configuration for the display
type Config struct {
}

// New returns a new ledstripdisplay driver given a LEDStriper, layout and rotation
func New(matrix *max72xx.Device) Device {
	return Device{
		matrix: matrix,
	}
}

// Configure initializes the display with default configuration
func (d *Device) Configure(cfg Config) {
	d.matrix.Configure()
	d.matrix.StopDisplayTest()
	d.matrix.SetDecodeMode(0)
	d.matrix.SetScanLimit(8)
	d.matrix.SetIntensity(8)
	d.matrix.StopShutdownMode()

}

// ClearDisplay erases the internal buffer
func (d *Device) ClearDisplay() {
	for i := 0; i < 8; i++ {
		d.buffer[i] = 0x00
	}
}

// SetPizel modifies the internal buffer.
func (d *Device) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return
	}
	if c.R != 0 || c.G != 0 || c.B != 0 {
		d.buffer[x] |= byte(0x80 >> (7 - y))
	} else {
		d.buffer[x] &^= byte(0x80 >> (7 - y))
	}
}

// Display sends the buffer (if any) to the screen.
func (d *Device) Display() error {
	for i := byte(1); i <= 8; i++ {
		d.matrix.WriteCommand(i, d.buffer[i-1])
	}
	return nil
}

// Size returns the current size of the display.
func (d *Device) Size() (w, h int16) {
	return 8, 8
}
