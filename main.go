package main

import (
	"fmt"
	stl "img_resizer/STL"
)

func main() {
	var prefix_b, prefix_s string = "b_", "s_"
	var maxWidthBigImg, maxHightBigImg uint = 1200, 900
	var maxWidthSmallImg, maxHightSmallImg uint = 186, 140

	imgPathes := stl.GetImgPathSlice(".")

	for _, imgPath := range imgPathes {
		fmt.Println(imgPath)
		err := stl.ResizeImg(maxWidthBigImg, maxHightBigImg, imgPath, prefix_b)
		if err != nil {
			panic(err)
		}
		err = stl.ResizeImg(maxWidthSmallImg, maxHightSmallImg, imgPath, prefix_s)
		if err != nil {
			panic(err)
		}
		err = stl.DeleteImg(imgPath)
		if err != nil {
			panic(err)
		}
	}
}
