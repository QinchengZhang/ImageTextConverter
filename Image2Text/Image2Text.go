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
)

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
			if r == 0 && g == 0 && b == 0 && a == 255 {
				break main
			}
			//fmt.Println(r, g, b, a)
			//r8, g8, b8, a8 := getRealRGBA(r, g, b, a)
			//fmt.Println(r8, g8, b8, a8)
			char := int32(r)<<24 | int32(g)<<16 | int32(b)<<8 | int32(a)
			//fmt.Printf("%c\n",char)
			res = append(res, rune(char))
		}
	}

	return string(res), nil
}

//func getRealRGBA(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {
//	if a == 0 {
//		return 0, 0, 0, 0
//	}
//	a >>= 8
//	return getRealValue(r, a), getRealValue(g, a), getRealValue(b, a), uint8(a)
//}
//func getRealValue(val, a uint32) uint8 {
//	val *= 0xff
//	val /= a
//	val >>= 8
//	return uint8(val)
//}

func getRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	str := fmt.Sprintf("%v", c)
	str = strings.TrimFunc(str, func(c rune) bool {
		return c == '{' || c == '}'
	})
	//fmt.Println(str)
	arr := strings.Split(str, " ")
	//fmt.Println(arr)
	r, _ := strconv.Atoi(arr[0])
	g, _ := strconv.Atoi(arr[1])
	b, _ := strconv.Atoi(arr[2])
	a, _ := strconv.Atoi(arr[3])
	return uint8(r), uint8(g), uint8(b), uint8(a) - 1
}
