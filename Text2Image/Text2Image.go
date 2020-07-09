package Text2Image

import (
	"bufio"
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
	"unicode/utf8"
)

func PathConvert2Image(textpath string, imagepath string) (err error) {
	// 注册异常处理函数
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("convert failed, please check the text file's encoding")
		}
	}()
	text, err := os.Open(textpath)
	if err != nil {
		return err
	}
	return Convert2Image(text, imagepath)
}

func Convert2Image(text *os.File, filepath string) (err error) {
	// 注册异常处理函数
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("convert failed, please check the text file's encoding")
		}
	}()
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

func Convert(text *os.File) (res image.Image, err error) {
	// 注册异常处理函数
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("convert failed, please check the text file's encoding")
		}
	}()
	info, err := text.Stat()
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(text)
	size := info.Size()
	if size == 0 {
		return nil, errors.New("do not convert an empty file")
	}
	width := int(math.Sqrt(float64(size)))
	height := int(info.Size())/width + int(size)%width
	x, y := 0, 0
	var img = image.NewNRGBA(image.Rect(x, y, width, height))
	for {
		data, _, err := r.ReadRune()
		rgb := make([]byte, 3)
		utf8.EncodeRune(rgb, data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		img.SetNRGBA(x, y, color.NRGBA{
			R: uint8(rgb[0]),
			G: uint8(rgb[1]),
			B: uint8(rgb[2]),
			A: 1,
		})
		if x += 1; x >= width {
			x = 0
			y += 1
		}
	}
	return img, nil
}
