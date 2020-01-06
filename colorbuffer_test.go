package colorbuffer

import (
	"testing"
)

const (
	width  = 20
	height = 20
)

func TestStride(t *testing.T) {
	cb := NewColorBuffer(width, height)

	if cb.Stride != width*4 {
		t.Error("Stride value is not set correctly")
	}
}

func TestPixOffset(t *testing.T) {
	var (
		cb       = NewColorBuffer(width, height)
		x        = 10
		y        = 20
		expected = cb.Stride*20 + x*4
		result   = cb.PixelOffset(x, y)
	)

	if result != expected {
		t.Errorf("Pixel offset is not as expected. Received: %d. Expected: %d", result, expected)
	}
}

func TestClear(t *testing.T) {
	cb := NewColorBuffer(width, height)

	white := uint32(0xFFFFFFFF)

	cb.Clear(white)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if cb.At(x, y) != white {
				t.Errorf("Image is not uniform white, 0x%-X", cb.At(x, y))
			}
		}
	}
}

func TestSet(t *testing.T) {
	cb := NewColorBuffer(width, height)

	cb.Set(10, 10, 0xFFFFFFFF)

	if cb.At(10, 10) != 0xFFFFFFFF {
		t.Error("Error setting the correct color")
	}
}

func TestAt(t *testing.T) {
	cb := NewColorBuffer(width, height)
	cb.Set(10, 10, 0xEEEEFFFF)

	if cb.At(10, 10) != 0xEEEEFFFF {
		t.Error("Error setting the correct color")
	}
}
