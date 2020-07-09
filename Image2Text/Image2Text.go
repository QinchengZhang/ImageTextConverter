/*
 * @Description:
 * @Autor: TJUZQC
 * @Date: 2020-07-06 14:21:33
 * @LastEditors: TJUZQC
 * @LastEditTime: 2020-07-06 14:22:01
 */
package image2text

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func PathConvert2Text(imagepath string, textpath string) error {
	image, err := os.Open(imagepath)
	if err != nil {
		return err
	}
	return Convert2Text(image, textpath)
}

func Convert2Text(image *os.File, filepath string) error {
	str, err := Convert(image)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}

func Convert(image *os.File) (string, error) {
	var res []rune
	data, err := png.Decode(image)
	if err != nil {
		return "", err
	}
	b := data.Bounds()
	width, height := b.Max.X, b.Max.Y
main:
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pix := data.At(x, y)
			r, g, b, a := getRGBA(pix)
			if a != 1 {
				break main
			}
			char, _ := utf8.DecodeRune([]byte{r, g, b})
			//fmt.Printf("%c\n",char)
			res = append(res, char)
		}
	}

	return string(res), nil
}

func getRGBA(c color.Color) (byte, byte, byte, byte) {
	str := fmt.Sprintf("%v", c)
	str = strings.TrimFunc(str, func(c rune) bool {
		return c == '{' || c == '}'
	})
	//fmt.Println(str)
	arr := strings.Split(str, " ")
	r, _ := strconv.Atoi(arr[0])
	g, _ := strconv.Atoi(arr[1])
	b, _ := strconv.Atoi(arr[2])
	a, _ := strconv.Atoi(arr[3])
	return byte(r), byte(g), byte(b), byte(a)
}
