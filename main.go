package main

import (
	"bufio"
	stl "img_resizer/STL"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	if len(os.Args) > 1 {
		commands := os.Args[1:]
		stl.Handle(commands)
	}

	var ImgPathes = stl.GetImgPathSlice(stl.PathRoot)

	for _, imgPath := range ImgPathes {
		err := stl.ResizeImg(stl.MaxWidthBigImg, stl.MaxHightBigImg, imgPath, stl.Prefix_b)
		if err != nil {
			stl.PrintError(err.Error())
			in.ReadString('\n')
			stl.Exit()
		}

		err = stl.ResizeImg(stl.MaxWidthSmallImg, stl.MaxHightSmallImg, imgPath, stl.Prefix_s)
		if err != nil {
			stl.PrintError(err.Error())
			in.ReadString('\n')
			stl.Exit()
		}

		err = stl.DeleteImg(imgPath)
		if err != nil {
			stl.PrintError(err.Error())
			in.ReadString('\n')
			stl.Exit()
		}

	}
	stl.PrintMessage("Процесс завершится после нажатия 'Enter'.")
	in.ReadString('\n')
	stl.Exit()
}
