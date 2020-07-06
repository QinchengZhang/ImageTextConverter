# ImageTextConverter
## 将文字转为图片

A text file(sunch as novel) convert to a png file like this

![test](.\img\test.png)

### example:

```go
package main

import (
	image2text "ImageTextConverter/Image2Text"
	"ImageTextConverter/Text2Image"
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("text/test.txt")
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}
	defer file.Close()
	err = Text2Image.Convert2Image(file, "img/test.png")
	if err != nil {
		fmt.Println("Image convert error: ", err)
		return
	}
}
```



Now you have a png file but cannot read,second step you can convert this image to a text

### example:

```go
package main

import (
	image2text "ImageTextConverter/Image2Text"
	"ImageTextConverter/Text2Image"
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("img/test.png")
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
}

```

