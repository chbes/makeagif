package frame

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"../ext"

	"github.com/nfnt/resize"
)

// Frame - Composant of GIF
type Frame struct {
	Data   image.Image
	Path   string
	Order  int
	Width  uint
	Height uint
}

// New - Create a frame, but data(image) is not load and generate
func New(p string, i int, w uint, h uint) Frame {
	return Frame{nil, p, i, w, h}
}

// Generate - Load input JPEG file (via path) and format in GIF
func (f *Frame) Generate() {
	err := f.loadImage()
	if err != nil {
		f.Data = nil
		return
	}

	f.resize()

	err = f.formatGIF()
	if err != nil {
		f.Data = nil
		return
	}

}

// NotGenerated - Return false is Data is not valid
func (f *Frame) NotGenerated() bool {
	return f.Data == nil
}

func (f *Frame) loadImage() error {
	fi, err := os.Open(f.Path)
	defer fi.Close()
	if err != nil {
		return err
	}

	if ext.IsJPEG(fi.Name()) {
		f.Data, err = jpeg.Decode(fi)
		if err != nil {
			return err
		}
	} else if ext.IsPNG(fi.Name()) {
		f.Data, err = png.Decode(fi)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *Frame) resize() {
	f.Data = resize.Resize(f.Width, f.Height, f.Data, resize.Lanczos3)
}

func (f *Frame) formatGIF() error {
	o := &gif.Options{}
	o.NumColors = 256
	buf := bytes.Buffer{}

	err := gif.Encode(&buf, f.Data, o)
	if err != nil {
		return err
	}

	f.Data, err = gif.Decode(&buf)
	return err
}
