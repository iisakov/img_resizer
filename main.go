package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	var maxWidth uint = 1200
	var maxHight uint = 900
	var newWidth, newHight uint

	imgFile, err := os.Open("test.jpg")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	newImgFile, err := os.Create("result.jpg")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	in := bufio.NewReader(imgFile)
	out := bufio.NewWriter(newImgFile)

	img, _, err := image.Decode(in)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bounds := img.Bounds()

	fmt.Println(bounds.Size())

	if uint(bounds.Size().X) > maxWidth {
		fmt.Println("Ширина изображения больше максимальной")

		newHight = uint(float32(maxWidth) / float32(bounds.Size().X) * float32(bounds.Size().Y))
		newWidth = maxWidth
		fmt.Printf("Новый размер изображения: %d %d \n", newWidth, newHight)

		if newHight > maxHight {
			fmt.Println("Высота изображения всё ещё больше максимальной")
			newWidth = uint(float32(bounds.Size().X) * float32(maxWidth) / float32(bounds.Size().Y))
			newHight = maxHight
			fmt.Printf("Новый размер изображения: %d %d \n", newWidth, newHight)
		}
	}

	newImg := resize.Resize(newWidth, newHight, img, resize.Lanczos3)

	// Encode uses a Writer, use a Buffer if you need the raw []byte
	err = jpeg.Encode(out, newImg, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// check err
}
