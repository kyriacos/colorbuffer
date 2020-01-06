package colorbuffer

// ColorBuffer holds all the information for the pixels.
type ColorBuffer struct {
	Pixels        []byte
	Stride        int // a.k.a the pitch (Width * 4).
	Width, Height int
}

// NewColorBuffer creates new ColorBuffer based on the supplied width and height.
// Sets the Stride at the time we initialize the ColorBuffer.
func NewColorBuffer(w, h int) *ColorBuffer {
	buf := make([]byte, w*h*4, w*h*4)
	return &ColorBuffer{
		Pixels: buf,
		Stride: w * 4,
		Width:  w,
		Height: h,
	}
}

// GetPitch returns the Stride.
// I just remember pitch instead of Stride so i added this method.
func (cb *ColorBuffer) GetPitch() int {
	return cb.Stride
}

// PixelOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).
func (cb *ColorBuffer) PixelOffset(x, y int) int {
	// Width*4*y + x*4
	return cb.Stride*y + x*4
}

// Set sets the values in the bytes buffer at x,y using the correct pixel offset.
func (cb *ColorBuffer) Set(x, y int, c uint32) {
	// cb[width*y*4+x*4+3] = byte(c)

	i := cb.PixelOffset(x, y)
	s := cb.Pixels[i : i+4 : i+4]
	s[0] = byte(c >> 24)
	s[1] = byte(c >> 16)
	s[2] = byte(c >> 8)
	s[3] = byte(c)
}

// At retrieves uint32 color value (ordered as rgba).
func (cb *ColorBuffer) At(x, y int) uint32 {
	offset := cb.PixelOffset(x, y)
	return uint32(cb.Pixels[offset])<<24 |
		uint32(cb.Pixels[offset+1])<<16 |
		uint32(cb.Pixels[offset+2])<<8 |
		uint32(cb.Pixels[offset+3])
}

// Clear clears the color buffer using the supplied color.
// If no color is supplied it defaults to Black (0x00000000).
func (cb *ColorBuffer) Clear(c ...uint32) {
	var col uint32 = 0x00000000

	if len(c) > 0 {
		col = c[0]
	}

	for x := 0; x < cb.Width; x++ {
		for y := 0; y < cb.Height; y++ {
			cb.Set(x, y, col)
		}
	}
}
