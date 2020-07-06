/*
 * @Description:
 * @Autor: TJUZQC
 * @Date: 2020-07-06 14:22:40
 * @LastEditors: TJUZQC
 * @LastEditTime: 2020-07-06 14:30:11
 */
package main

import (
	image2text "ImageTextConverter/Image2Text"
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("img/llzjpbm.png")
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}
	defer file.Close()
	err = image2text.Convert2Text(file, "text/new.txt")
	if err != nil {
		fmt.Println("Image convert error: ", err)
		return
	}

	//var file, err = os.Open("text/六零之极品爸妈.txt")
	//if err != nil {
	//	fmt.Println("File open error: ", err)
	//	return
	//}
	//defer file.Close()
	//err = Text2Image.Convert2Image(file, "img/llzjpbm.png")
	//if err != nil {
	//	fmt.Println("Image convert error: ", err)
	//	return
	//}
}
