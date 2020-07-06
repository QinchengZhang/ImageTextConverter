package Text2Image

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
)

func Convert2Image(text *os.File, filepath string) error {
	img, err := Convert(text)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	//fmt.Println(img)
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

func Convert(text *os.File) (image.Image, error) {
	info, err := text.Stat()
	if err != nil {
		return &image.NRGBA{}, err
	}
	r := bufio.NewReader(text)
	size := info.Size()
	width := int(math.Sqrt(float64(size)))
	height := int(info.Size())/width + int(size)%width
	x, y := 0, 0
	var img = image.NewNRGBA(image.Rect(x, y, width, height))
	for {
		data, _, err := r.ReadRune()
		//fmt.Println(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return &image.NRGBA{}, err
		}
		img.SetNRGBA(x, y, color.NRGBA{
			R: uint8(data >> 24),
			G: uint8(data >> 16),
			B: uint8(data >> 8),
			A: uint8(data) + 1,
		})
		if x += 1; x >= width {
			x = 0
			y += 1
		}
		//fmt.Println(uint8(data >> 16), uint8(data >> 8), uint8(data), 1)
	}
	return img, nil
}
