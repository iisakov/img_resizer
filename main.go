package main

import (
	"bufio"
	"fmt"
	stl "img_resizer/STL"
	"os"
)

func main() {
	var prefix_b, prefix_s string = "b_", "s_"
	var maxWidthBigImg, maxHightBigImg uint = 1200, 900
	var maxWidthSmallImg, maxHightSmallImg uint = 186, 140
	in := bufio.NewReader(os.Stdin)

	imgPathes := stl.GetImgPathSlice(".")

	for _, imgPath := range imgPathes {
		err := stl.ResizeImg(maxWidthBigImg, maxHightBigImg, imgPath, prefix_b)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Print("Процесс завершится после нажатия 'Enter'.")
			in.ReadString('\n')
			panic(err)
		}
		fmt.Println(maxWidthBigImg, maxHightBigImg, imgPath, prefix_b)

		err = stl.ResizeImg(maxWidthSmallImg, maxHightSmallImg, imgPath, prefix_s)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Print("Процесс завершится после нажатия 'Enter'.")
			in.ReadString('\n')
			panic(err)
		}
		fmt.Println(maxWidthBigImg, maxHightBigImg, imgPath, prefix_s)

		err = stl.DeleteImg(imgPath)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Print("Процесс завершится после нажатия 'Enter'.")
			in.ReadString('\n')
			panic(err)
		}
		fmt.Println(imgPath, "Удалён.\n")

	}
	fmt.Print("Процесс завершится после нажатия 'Enter'.")
	in.ReadString('\n')
}
